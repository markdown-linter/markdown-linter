# Markdown Linter

![Go](https://github.com/markdown-linter/markdown-linter/workflows/Go/badge.svg?branch=master)

## Usage

By default markdown-linter prints help:

```bash
markdown-linter
```

Result:

```bash
Usage: markdown-linter [--version] [--help] <command> [<args>]

Available commands are:
    generate-config    Generates configuration file from the template
    lint               Lints .md files
```

### Generate config

The next command will create `.markdown-linter.yml` in the current directory:

```bash
markdown-linter create-config
```

If file exists the command will exit. But you can override this behavior and
overwrite existed config with `-force` flag.

### Lint

Examples:

```bash
markdown-linter lint -R . -D testdata -f README.md
```
>>>>>>> c712430... Reorganize files
