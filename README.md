# Glia Regexp Utility

Glia Regex Utility is a small executable that is meant to accomplish two things:
- Mask PII in chat messages using regular expressions provided by clients
- Validate regular expressions provided by clients

This executable is meant to be a very lightweight, is using only Golang's standard
library and has zero external dependencies.

## Usage

Sub-commands and their flags can be seen by executing one of the following commands.

```
./regexp_utility --help
./regexp_utility mask --help
./regexp_utility validate --help
```

### Input masking

```bash
./regexp_utility mask \
    --input "My SSN is 111-12-1234 and my code is secret" \
    --regexp "secret" \
    --regexp "(?:00[1-9]|0[1-9][0-9]|[1-578][0-9]{2}|6[0-57-9][0-9]|66[0-57-9])-(?:0[1-9]|[1-9]0|[1-9][1-9])-(?:[1-9][0-9][0-9][0-9]|[0-9][1-9][0-9][0-9]|[0-9][0-9][1-9][0-9]|[0-9][0-9][0-9][1-9])"
```

### Validating regex

Valid regex example:

```bash
./regexp_utility validate --regexp "(abc|def)"
```

Invalid regex example (does not conform to RE2 syntax):

```bash
./regexp_utility validate --regexp '^((?!666|000)[0-8][0-9\_]{2}\-(?!00)[0-9\_]{2}\-(?!0000)[0-9\_]{4})*$'
```

## Development

While developing the executable it is not necessary to build the executable every time. It is possible to use `go run`
command with all the possible flags. For example:

```bash
go run regexp_utility.go validate --regexp "foo"
```

### Prerequisites

To build the executable you need to have Go installed locally. This can be done by one of the following methods.
There is no preference which method to use.

#### Installing go using asdf

NB! At the time of last update of this document the latest version of Go was `1.18`.
Go versions are usually backwards compatible and it will most likely work with the newer version.

`asdf plugin-add golang && asdf install golang 1.18 && asdf global golang 1.18`

#### Installing go using Homebrew

`brew install go`

### Running tests

`go test ./...` or `go test -v ./...` for verbose output

### Building the executable

To build executable on Mac the following command needs to be run in the project directory.

```bash
go build regex_utility.go
```

It is also possible to build an executable that can be run on Linux hosts. This can be useful to test the changes in
DevSpaces with the new version of the executable. To build an executable for linux run the following command.

```bash
GOOS=linux GOARCH=amd64 go build regex_utility.go
```