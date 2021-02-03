package main

import (
	"fmt"

	"github.com/youngjuning/go-release"
)

func main() {
	update, err := release.CheckUpdate("youngjuning", "tpc", "0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Println(update.IsUpdate)
}
