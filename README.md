# printbuildinfo

A small example Go program that prints some of the information returned by
[`debug.ReadBuildInfo`](https://pkg.go.dev/runtime/debug#ReadBuildInfo).  The goal of the example is to show how a
program can know its own version tag when it is `go install`ed.

## Usage

```shell
go install github.com/linux2647/printbuildinfo@latest
printbuildinfo
```