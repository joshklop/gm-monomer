package main

import (
	"os"

	"github.com/joshklop/monomer-cosmos-sdk/server"
	svrcmd "github.com/joshklop/monomer-cosmos-sdk/server/cmd"

	"github.com/joshklop/gm-monomer/app"
	"github.com/joshklop/gm-monomer/cmd/gmd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
