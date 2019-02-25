# Go module example

Basic example to try out how `go mod` works.

```bash

Go version 1.11.5

```

## Steps to recreate

Create a directory outside of `GOPATH`. In this example, it is done by `mkdir go-mod-example`.

Enable go modules by running `go mod init go-mod-example` which is the name of the repository. It will create a file called `go.mod` in the main root directory.

Do coding as usually by import packages in `main.go`.

Run `go build`. It would automatically downloads packages require in the code and appends new packages in `go.mod` and will create a file calls `go.sum`.

## After cloning this repo

Run `go build`.
