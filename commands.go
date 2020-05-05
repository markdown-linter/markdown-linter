package main

import (
	"github.com/markdown-linter/markdown-linter/command"
	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"generate-config": func() (cli.Command, error) {
			return &command.GenerateConfigCommand{}, nil
		},
		"lint": func() (cli.Command, error) {
			return &command.LintCommand{}, nil
		},
	}
}
