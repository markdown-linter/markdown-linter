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

Result:

```bash
+------------------------------+------+--------------+--------------------------------+
|             FILE             | LINE |    PLUGIN    |          DESCRIPTION           |
+------------------------------+------+--------------+--------------------------------+
| testdata/another/test.md     |    1 | MissingH1Tag | Empty file could not be linted |
| testdata/another.md          |    1 | MissingH1Tag | Empty file could not be linted |
| testdata/markdown.md/test.md |    1 | MissingH1Tag | Empty file could not be linted |
| testdata/test.md             |    3 | FixmeTag     | The line has FIXME tag         |
+------------------------------+------+--------------+--------------------------------+
```
