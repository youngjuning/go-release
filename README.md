# go-release

A version control tools based on github release.

## Why Use?

We assume the following conditions:

1. You have a binary executable and publish it to GitHub release like [Deno](https://github.com/denoland/deno/releases) do.
2. You want to realize `upgrade` command like `deno upgrade`

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

## Use in Cobra

```go
package main

import (
	"fmt"

	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cobra"
	"github.com/youngjuning/go-release"
)

const Version = "0.0.1"

func checkUpgrade(current string, force bool) {
	if force {
		fmt.Println("Looking up latest version")
	}
	update, err := release.CheckUpdate("youngjuning", "tpc", current)
	if err != nil {
		panic(err)
	}
	if update.IsUpdate {
		if force {
			fmt.Printf("Found latest version %v \n", update.LatestVersion)
		} else {
			fmt.Printf("Found tpc latest version %v \n", update.LatestVersion)
		}
		sh.Command("bash", "-c", "curl -fsSL https://raw.githubusercontent.com/youngjuning/tpc/main/install.sh | sh").Run()
		if !force {
			fmt.Println("\nPress any key to exit.")
		}
	} else {
		if force {
			fmt.Printf("Local tpc version %v is the most recent release \n", current)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:     "tpc",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		sh.Command("tpc", "-h").Run()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		sh.Command("bash", "-c", "tpc upgrade --force=false").Start()
	},
}

var cmdUpgrade = &cobra.Command{
	Use: "upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		checkUpgrade(Version, force)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {},
}

func main() {
	cmdUpgrade.Flags().Bool("force", true, "Force to upgrade")
	rootCmd.AddCommand(cmdUpgrade)
	rootCmd.Execute()
}
```

## TODO

- [ ] Add RunInstaller function to install binary executable.
- [ ] Add Test (I can't write it. SOS)

## Thanks

This project is inspired on https://github.com/denoland/deno. Thanks for the author.
