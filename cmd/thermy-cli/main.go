package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

const appUrl = "http://localhost:8080"
const apiPrefix = "/api/v1"

var roles = []string{"student", "educator", "admin"}
var unitFilters = []string{"all", "models", "properties"}
var propertyFilters = []string{"all", "unit"}

func main() {
	var username, password, token, layer, filter string
	filenamePtr := new(string)

	defaultQueryParams := map[string]*string{
		"token": &token,
		"layer": &layer,
	}

	tokenQueryParam := map[string]*string{
		"token": &token,
	}

	app := &cli.App{
		Name:    "thermy-cli",
		Usage:   "Console UI for thermy backend",
		Version: "v0.2.0",
		Commands: []*cli.Command{
			{
				Name:  "login",
				Usage: "login to the app",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "user",
						Aliases:     []string{"u"},
						Usage:       "specify a `USERNAME`",
						Destination: &username,
						Action:      checkStringFlag("username"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "password",
						Aliases:     []string{"p"},
						Usage:       "specify a `PASSWORD`",
						Destination: &password,
						Action:      checkStringFlag("password"),
						Required:    true,
					},
				},
				Action: login("/login", &username, &password),
			},
			{
				Name:  "logout",
				Usage: "logout from the app",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
				},
				Action: logout("/logout", &token),
			},
			{
				Name:  "stat",
				Usage: "get information about app",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
				},
				Action: stat("/admin/stat", &token),
			},
			{
				Name:  "layers",
				Usage: "commands for working with layers",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "show all text markup layers",
						Action: commonHandler(http.MethodGet, "/layers/all", nil, tokenQueryParam, nil),
					},
					{
						Name:  "add",
						Usage: "add text markup layer",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "layer",
								Aliases:     []string{"l"},
								Usage:       "specify a text markup `LAYER`",
								Destination: &layer,
								Action:      checkStringFlag("layer"),
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodPost, "/layers", nil, defaultQueryParams, nil),
					},
				},
			},
			{
				Name:  "models",
				Usage: "commands for working with structural models",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "layer",
						Aliases:     []string{"l"},
						Usage:       "specify a text markup `LAYER`",
						Destination: &layer,
						Action:      checkStringFlag("layer"),
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "show a list of models",
						Action: commonHandler(http.MethodGet, "/models/all", nil, defaultQueryParams, nil),
					},
					{
						Name:  "add",
						Usage: "add models",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file",
								Aliases:     []string{"f"},
								Usage:       "Load request body from input `FILE`",
								Destination: filenamePtr,
								Action:      checkJSON,
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodPost, "/models", nil, defaultQueryParams, filenamePtr),
					},
				},
			},
			{
				Name:  "elements",
				Usage: "commands for working with elements of structural models",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "layer",
						Aliases:     []string{"l"},
						Usage:       "specify a text markup `LAYER`",
						Destination: &layer,
						Action:      checkStringFlag("layer"),
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "show a list of model elements",
						Action: commonHandler(http.MethodGet, "/elements/all", nil, defaultQueryParams, nil),
					},
					{
						Name:  "add",
						Usage: "add models",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file",
								Aliases:     []string{"f"},
								Usage:       "Load request body from input `FILE`",
								Destination: filenamePtr,
								Action:      checkJSON,
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodPost, "/elements", nil, defaultQueryParams, filenamePtr),
					},
				},
			},
			{
				Name:  "properties",
				Usage: "commands for working with properties",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						Usage:       "Load request body from input `FILE`",
						Destination: filenamePtr,
						Action:      checkJSON,
						Required:    false,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "show a list of properties according to the filter",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "filter",
								Usage:       fmt.Sprintf("specify a filter `OPTION` for searching properties (one of %v)", propertyFilters),
								Destination: &filter,
								Action:      checkFilterFlag(propertyFilters),
								Required:    true,
							},
							&cli.StringFlag{
								Name:        "layer",
								Aliases:     []string{"l"},
								Usage:       "specify a text markup `LAYER`",
								Destination: &layer,
								Action:      checkStringFlag("layer"),
								Required:    false,
							},
						},
						Action: commonHandler(http.MethodPut, "/properties", &filter, defaultQueryParams, filenamePtr),
					},
					{
						Name:   "add",
						Usage:  "add properties",
						Action: commonHandler(http.MethodPost, "/properties", nil, tokenQueryParam, filenamePtr),
					},
				},
			},
			{
				Name:  "units",
				Usage: "commands for working with units",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "layer",
						Aliases:     []string{"l"},
						Usage:       "specify a text markup `LAYER`",
						Destination: &layer,
						Action:      checkStringFlag("layer"),
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						Usage:       "Load request body from input `FILE`",
						Destination: filenamePtr,
						Action:      checkJSON,
						Required:    false,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "show a list of units according to the filter",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "filter",
								Usage:       fmt.Sprintf("specify a filter `OPTION` for searching units (one of %v)", unitFilters),
								Destination: &filter,
								Action:      checkFilterFlag(unitFilters),
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodPut, "/units", &filter, defaultQueryParams, filenamePtr),
					},
					{
						Name:   "add",
						Usage:  "add units",
						Action: commonHandler(http.MethodPost, "/units", nil, defaultQueryParams, filenamePtr),
					},
					{
						Name:   "edit",
						Usage:  "edit units",
						Action: commonHandler(http.MethodPatch, "/units", nil, defaultQueryParams, filenamePtr),
					},
				},
			},
			{
				Name:  "users",
				Usage: "commands for working with users",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{"t"},
						Usage:       "specify an authentication `TOKEN` that is issued at login",
						Destination: &token,
						Action:      checkStringFlag("token"),
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add user",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file",
								Aliases:     []string{"f"},
								Usage:       "Load request body from input `FILE`",
								Destination: filenamePtr,
								Action:      checkJSON,
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodPost, "/users", nil, tokenQueryParam, filenamePtr),
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
