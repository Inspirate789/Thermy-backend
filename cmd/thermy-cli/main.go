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

var roles = []string{"student", "educator", "admin"}
var unitFilters = []string{"all", "models", "properties"}

func main() {
	var username, password, token, layer, filter, role string
	var filenamePtr *string

	defaultQueryParams := map[string]*string{
		"token": &token,
		"layer": &layer,
	}

	app := &cli.App{
		Name:    "thermy-cli",
		Usage:   "Technological UI for thermy backend",
		Version: "v0.1.0",
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
				Name:  "unit",
				Usage: "options for working with units",
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
						Name:        "role",
						Aliases:     []string{"r"},
						Usage:       fmt.Sprintf("specify a user `ROLE` (one of %v)", roles),
						Destination: &role,
						Action:      checkRoleFlag(roles),
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
						Name:  "search",
						Usage: "search units",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "filter",
								Usage:       fmt.Sprintf("specify a filter `OPTION` for searching units (one of %v)", unitFilters),
								Destination: &filter,
								Action:      checkFilterFlag(unitFilters),
								Required:    true,
							},
						},
						Action: commonHandler(http.MethodGet, "/units", &role, &filter, defaultQueryParams, filenamePtr),
					},
					{
						Name:   "store",
						Usage:  "store units",
						Action: commonHandler(http.MethodPost, "/units", &role, nil, defaultQueryParams, filenamePtr),
					},
					{
						Name:   "edit",
						Usage:  "edit units",
						Action: commonHandler(http.MethodPatch, "/units", &role, nil, defaultQueryParams, filenamePtr),
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
