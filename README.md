# goswitch

go version switching utility.
goswitch is just a wrapper for `go get` and `ln -s`.

## How to use

### Prerequisites
- Go 1.13+

### Installing
```
$ go get -u github.com/akito0107/goswitch
```

### Example

1. check current version.

```
$ go version
go version go1.13.10 darwin/amd64
```

2. show all available versions by `goswitch ls-remote`

```
$ goswitch ls-remote

available versions:

go1
go1.2.2
go1.3, go1.3.1, go1.3.2, go1.3.3
go1.4, go1.4.1, go1.4.2, go1.4.3
go1.5, go1.5.1, go1.5.2, go1.5.3, go1.5.4
go1.6, go1.6.1, go1.6.2, go1.6.3, go1.6.4
go1.7, go1.7.1, go1.7.3, go1.7.4, go1.7.5, go1.7.6
go1.8, go1.8.1, go1.8.2, go1.8.3, go1.8.4, go1.8.5, go1.8.6, go1.8.7
go1.9, go1.9.1, go1.9.2, go1.9.3, go1.9.4, go1.9.5, go1.9.6, go1.9.7
go1.10, go1.10.1, go1.10.2, go1.10.3, go1.10.4, go1.10.5, go1.10.6, go1.10.7, go1.10.8
go1.11, go1.11.1, go1.11.2, go1.11.3, go1.11.4, go1.11.5, go1.11.6, go1.11.7, go1.11.8, go1.11.9, go1.11.10, go1.11.11, go1.11.12, go1.11.13
go1.12, go1.12.1, go1.12.2, go1.12.3, go1.12.4, go1.12.5, go1.12.6, go1.12.7, go1.12.8, go1.12.9, go1.12.10, go1.12.11, go1.12.12, go1.12.13, go1.12.14, go1.12.15, go1.12.16, go1.12.17
go1.13, go1.13.1, go1.13.2, go1.13.3, go1.13.4, go1.13.5, go1.13.6, go1.13.7, go1.13.8, go1.13.9, go1.13.10
go1.14, go1.14.1, go1.14.2
```

3. switch go version by `goswitch use`

```
$ goswitch use go1.14.2
2020/04/09 11:22:24 start go get...
2020/04/09 11:22:24 go get golang.org/dl/go1.14.2
go: downloading golang.org/dl v0.0.0-20200408221700-d6f4cf58dce2
go: found golang.org/dl/go1.14.2 in golang.org/dl v0.0.0-20200408221700-d6f4cf58dce2
2020/04/09 11:22:27 go get finished. start download
Downloaded   0.0% (    14448 / 125040726 bytes) ...
....
Downloaded 100.0% (125040726 / 125040726 bytes)
Unpacking /Users/akito.ito/sdk/go1.14.2/go1.14.2.darwin-amd64.tar.gz ...
Success. You may now run 'go1.14.2'
2020/04/09 11:23:23 download finished.
2020/04/09 11:23:23 switch go version
```

4. check current version 

```
$ go version
go version go1.14.2 darwin/amd64
```

## Options
```sh
$ goswitch -h
NAME:
   goswitch - go version switching utility

USAGE:
   goswitch [global options] command [command options] [arguments...]

COMMANDS:
   use        switch current go version
   ls-remote  show all available versions
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## License
This project is licensed under the Apache License 2.0 License - see the [LICENSE](LICENSE) file for details
