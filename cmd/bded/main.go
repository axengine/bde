package main

import (
	"os"

	"github.com/axengine/bde/cmd/bded/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/axengine/bde/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
