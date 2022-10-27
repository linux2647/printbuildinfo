package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintf(os.Stderr, "An error occurred reading build info\n")
	}
	fmt.Println(buildInfo.Main.Path, buildInfo.Main.Version)

	fmt.Println("Settings:")
	for _, setting := range buildInfo.Settings {
		fmt.Println(" ", setting.Key, "=", setting.Value)
	}
}
