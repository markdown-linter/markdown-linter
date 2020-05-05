package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type stringList []string

func (a *stringList) String() string {
	return ""
}

func (a *stringList) Set(value string) error {
	*a = append(*a, value)

	return nil
}

func parseFlags(command string, args []string, fn func(*flag.FlagSet)) *flag.FlagSet {
	f := flag.NewFlagSet(command, flag.ExitOnError)

	fn(f)

	_ = f.Parse(args)

	return f
}

func GetMarkdownFilesInDirectory(path string) ([]string, error) {
	files := make([]string, 0)

	if err := checkDirectory(path); err != nil {
		return files, err
	}

	glob, err := filepath.Glob(path + string(os.PathSeparator) + "*.md")

	if err != nil {
		return files, fmt.Errorf("unable to get files from directory: %s", err.Error())
	}

	for _, path := range glob {
		if stat, _ := os.Stat(path); stat.IsDir() {
			continue
		}

		files = append(files, path)
	}

	return files, nil
}

func GetMarkdownFilesInDirectoryRecursively(path string) ([]string, error) {
	var err error

	files := make([]string, 0)

	if err = checkDirectory(path); err != nil {
		return files, err
	}

	if files, err = walkAllMarkdownFilesInDir(path); err != nil {
		return files, err
	}

	return files, nil
}

func walkAllMarkdownFilesInDir(path string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.Mode().IsRegular() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}

func checkDirectory(path string) error {
	stat, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("directory %q does not exist", path)
	}

	if !stat.IsDir() {
		return fmt.Errorf("%q is not a directory", path)
	}

	return nil
}
