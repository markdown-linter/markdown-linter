package main

import (
	"log"
	"os"
	"runtime"

	"github.com/mitchellh/cli"
)

func main() {
	log.SetOutput(os.Stderr)

	log.Printf("[INFO] markdown-linter version: %s %s", Version, GitCommit)
	log.Printf("[INFO] Go runtime version: %s", runtime.Version())
	log.Printf("[INFO] CLI args: %-v", os.Args)

	cliRunner := &cli.CLI{
		Name:     "markdown-linter",
		Version:  Version,
		Args:     os.Args[1:],
		Commands: Commands(),
	}

	exitCode, err := cliRunner.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitCode)
}
