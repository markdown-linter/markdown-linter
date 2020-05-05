package command

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type GenerateConfigCommand struct{}

type configFlags struct {
	configPath string
	force      bool
}

func (c *GenerateConfigCommand) Run(args []string) int {
	flags := c.parseFlags(args)

	if !c.isFileCreatable(flags.configPath, flags.force) {
		return 1
	}

	return c.generateConfig(flags.configPath)
}

func (c *GenerateConfigCommand) parseFlags(args []string) configFlags {
	var flags configFlags

	_ = parseFlags("generate-config", args, func(f *flag.FlagSet) {
		f.StringVar(&flags.configPath, "config", ".markdown-linter.yml", "Path to config file")
		f.BoolVar(&flags.force, "force", false, "Use to overwrite existed config")
	})

	return flags
}

func (c *GenerateConfigCommand) isFileCreatable(filePath string, force bool) bool {
	info, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return true
		}

		log.Printf("Unhandled error happened: %s", err.Error())

		return false
	}

	if info.IsDir() {
		log.Printf("%q is a directory", filePath)

		return false
	}

	if force {
		return true
	}

	log.Printf("File %q already exists. If you want to overwrite it use -force flag", filePath)

	return false
}

func (c *GenerateConfigCommand) Help() string {
	helpText := `
Usage: markdown-linter generate-config [options]

Options:

  -config string   Path to config file (default ".markdown-linter.yml")
  -force           Use to overwrite existed config

Examples:

  markdown-linter generate-config -config ~/.markdown-linter.yml -force
`

	return strings.TrimSpace(helpText)
}

func (c *GenerateConfigCommand) Synopsis() string {
	return "Generates configuration file from the template"
}

func (c *GenerateConfigCommand) generateConfig(configPath string) int {
	content := []byte(c.configContent())

	err := ioutil.WriteFile(configPath, content, 0600)

	if err != nil {
		log.Printf("Unable to write to file %q. Error: %s", configPath, err.Error())

		return 1
	}

	return 0
}

func (c *GenerateConfigCommand) configContent() string {
	return "plugins: []\n"
}
