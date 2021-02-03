package main

import (
	"fmt"

	"github.com/youngjuning/go-release"
)

func main() {
	update, err := release.CheckUpdate("denoland", "deno", "0.0.1")
	if err != nil {
		panic(err)
	}
	if update.IsUpdate {
		fmt.Printf("Latest version is %v.\n", update.LatestVersion)
	}
}
