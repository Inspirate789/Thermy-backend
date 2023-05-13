package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/buger/jsonparser"
	"github.com/urfave/cli/v2"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func login(path string, username, password *string) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		req, err := createRequest(http.MethodPost, path, url.Values{}, &entities.AuthRequest{
			Username: *username,
			Password: *password,
		})
		if err != nil {
			return err
		}

		resBody, err := doRequest(req)
		if err != nil {
			return err
		}
		token, err := jsonparser.GetString(resBody, "token")
		if err != nil {
			return fmt.Errorf("client: error extracting `token` from response body: %w", err)
		}
		fmt.Printf("Authentication token: %s\n", token)

		return nil
	}
}

func logout(path string, token *string) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		reqArgs := url.Values{}
		reqArgs.Set("token", *token)
		req, err := createRequest(http.MethodPost, path, reqArgs, nil)
		if err != nil {
			return err
		}

		_, err = doRequest(req)

		return err
	}
}

func stat(path string, token *string) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		reqArgs := url.Values{}
		reqArgs.Set("token", *token)
		req, err := createRequest(http.MethodGet, path, reqArgs, nil)
		if err != nil {
			return err
		}

		resBody, err := doRequest(req)
		if err != nil {
			return err
		}

		prettyJSON := new(bytes.Buffer)
		err = json.Indent(prettyJSON, resBody, "", "\t")
		if err != nil {
			return fmt.Errorf("client: error parsing the response body: %w", err)
		}
		fmt.Println(prettyJSON.String())

		return nil
	}
}

func commonHandler(method, path string, option *string, queryParams map[string]*string, filenamePtr *string) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if option != nil {
			path += "/" + *option
		}

		var body []byte
		if filenamePtr != nil && *filenamePtr != "" {
			var err error
			body, err = os.ReadFile(*filenamePtr)
			if err != nil {
				return fmt.Errorf("could not read file %s: %w", *filenamePtr, err)
			}
		}
		values := url.Values{}
		for key, value := range queryParams {
			values.Set(key, *value)
		}
		req, err := createRequest(method, path, values, body)
		if err != nil {
			return err
		}

		//reqDump, err := httputil.DumpRequestOut(req, true)
		//if err != nil {
		//	return err
		//}
		//
		//fmt.Printf("REQUEST:\n%s", string(reqDump))

		resBody, err := doRequest(req)
		if err != nil {
			return err
		}

		if len(resBody) > 0 {
			err = jsonparser.ObjectEach(resBody, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				k := strings.Replace(string(key), "_", " ", -1) + ": "
				fmt.Printf("%s\n%s\n", k, string(value))
				return nil
			})
			if err != nil {
				return fmt.Errorf("client: error iterating the response body: %w", err)
			}
		}

		return nil
	}
}
