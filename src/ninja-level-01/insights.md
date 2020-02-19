# Highlights

Installation and setup of directory for Go.

## Directory Setpup

By default go expects the go binary and workspace added in
to the path. So there should be a GoProject folders with src, bin and pkg folder in it. Then GOROOT, GOBIN need to be setted up.

```bash
export GOPATH="$HOME/GoProjects"
export GOBIN="$GOPATH/bin"

# depends on the operating system
export GOROOT=/usr/local/go/
export PATH=$PATH:$GOROOT/bin
```

`go env` commadn can be used to check the current setup.
