# go-release

A version control tools based on github release.

## Install

```sh
$ go get github.com/youngjuning/go-release
```

## Example

```go
package main

import "github.com/youngjuning/go-release"

func main() {
  update, err := release.CheckUpdate("youngjuning", "tpc", "0.0.1")
  if err != nil {
    panic(err)
  }
  if update.IsUpdate {
    fmt.Printf("Latest version is %v.\n",update.LatestVersion)
  }
}
```

## TODO

- [ ] Add RunInstaller function to install binary executable.
- [ ] Add Test (I can't write it. SOS)

## Thanks

This project is inspired on https://github.com/denoland/deno. Thanks for the author.
