# mdrun

mdrun is task runnner of shell commands.
You can select running commands with fizzy finder from sample commands written in README.md or another specified Markdown file.

## Setup

```
go install github.com/yammerjp/mdrun
```

## Run

### List Commands

```
$ mdrun list
```

### Show a Command

```
$ mdrun --dry-run
```

### Run a Command

```
$ mdrun
```

### Help

```
$ mdrun --help
```

### Specify Task Defined Markdown File

```
$ mdrun --target path/to/markdownfile.md
```

## Development

### test

```
$ go test ./...
```
