# go-release

A version control tools based on github release.

## Why Use?

We assume the following conditions:

1. We have a binary executable and publish it on GitHub release like [Deno](https://github.com/denoland/deno/releases) do.
2. We want to realize `upgrade` command like `deno upgrade`

> Try me help you to do that.

## Install

```sh
$ go get github.com/youngjuning/go-release
```

## Example

```go
package main

import "github.com/youngjuning/go-release"

func main() {
  update, err := release.CheckUpdate("denoland", "deno", "0.0.1")
  if err != nil {
    panic(err)
  }
  if update.IsUpdate {
    fmt.Printf("Latest version is %v.\n",update.LatestVersion) // out: Latest version is 1.7.1.
  }
}
```

## TODO

- [ ] Add RunInstaller function to install binary executable.
- [ ] Add Test (I can't write it. SOS)

## Thanks

This project is inspired on https://github.com/denoland/deno. Thanks for the author.
