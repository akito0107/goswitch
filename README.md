# goswitch

go version switching utility.
goswitch is just a wrapper for `go get` and `ln -s`.

## Getting Started

### Prerequisites
- Go 1.13+

### Workflow

`goswitch` makes symlink for `go` executable command under `$GOBIN` or `$GOPATH/bin`.

### Installing
```
$ go get -u github.com/akito0107/goswitch
```

## Options
```sh
$ goswitch -h
NAME:
   goswitch - go version switching utility

USAGE:
   goswitch [global options] command [command options] [arguments...]

COMMANDS:
   use
   ls-remote
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## License
This project is licensed under the Apache License 2.0 License - see the [LICENSE](LICENSE) file for details
