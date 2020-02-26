# Highlights

Installation and setup of directory for Go.

## Directory Setpup

By default go expects the go binary and workspace added in to the path. So there should be a GoProject folders with src, bin and pkg folder in it. Then GOROOT, GOBIN need to be setted up.

```bash
export GOPATH="$HOME/GoProjects"
export GOBIN="$GOPATH/bin"

# depends on the operating system
export GOROOT=/usr/local/go/
export PATH=$PATH:$GOROOT/bin
```

`go env` commadn can be used to check the current setup.

## Package Documentation

golang.org is official go language and package information docuemntaiton, Where godoc.org includes documentation of 3rd party and official packages for golang.

## Go Package Manage with Go Module

Go module help to set up the workspace generally to work anywhere in the machine. It also allows to lockdown the packages. In go software dependencies is considered as first class concept. Detail can be found in using go module
[doc](https://blog.golang.org/using-go-modules).

## Referance

1. [Quick Lookup Doc](https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit)
