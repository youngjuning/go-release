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
	// Install to ~/.tuya/bin/tpc
	if update.IsUpdate {
		if force {
			fmt.Printf("Found latest version %v \n", update.LatestVersion)
		} else {
			fmt.Printf("Found tpc latest version %v \n", update.LatestVersion)
		}
		release.InstallLatest(update.LatestReleaseURL, "tpc", ".tuya")
		if !force {
			fmt.Println("\nPress any key to exit.")
		}
	} else {
		if force {
			fmt.Printf("Local tpc version %v is the most recent release \n", current)
		}
	}
}

// rootCmd 代表没有调用子命令时的基础命令
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
	// 不要编辑，只是用来覆盖 rootCmd 的 PersistentPostRun
	PersistentPostRun: func(cmd *cobra.Command, args []string) {},
}

func main() {
	cmdUpgrade.Flags().Bool("force", true, "Force to upgrade")
	rootCmd.AddCommand(cmdUpgrade)
	// 初始化应用
	rootCmd.Execute()
}
