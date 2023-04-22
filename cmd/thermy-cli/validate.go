package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func checkStringFlag(flagName string) func(*cli.Context, string) error {
	return func(context *cli.Context, s string) error {
		if s == "" {
			return fmt.Errorf("the use of an empty %s is forbidden", flagName)
		}
		return nil
	}
}

func checkRoleFlag(roles []string) func(*cli.Context, string) error {
	return func(context *cli.Context, s string) error {
		for _, filter := range roles {
			if s == filter {
				return nil
			}
		}
		return fmt.Errorf("role %s is undefined", s)
	}
}

func checkFilterFlag(filters []string) func(*cli.Context, string) error {
	return func(context *cli.Context, s string) error {
		for _, filter := range filters {
			if s == filter {
				return nil
			}
		}
		return fmt.Errorf("filter option %s is undefined", s)
	}
}

func checkJSON(cCtx *cli.Context, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", filename, err)
	}

	if !json.Valid(data) {
		return fmt.Errorf("file %s is invalid", filename)
	}

	return nil
}
