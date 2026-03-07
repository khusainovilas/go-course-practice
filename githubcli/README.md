# GitHub CLI

CLI tool in Go to get information about GitHub repositories.

## Install

```bash
go build -o githubcli.exe .
```

Note: On Linux/Mac use `githubcli` instead of `githubcli.exe`.

## Usage

```bash
githubcli <url>
```

## Examples

```bash
# All data
githubcli.exe https://github.com/mimi-net/miminet

# With flags
githubcli.exe https://github.com/mimi-net/miminet -name -star
```

## Flags

-name - repository name
-desc - description  
-star - stars
-forks - forks
-date - creation date

## Requirements

Go 1.25+
