package main

import (
	"fmt"
	"os"

	svrcmd "github.com/joshklop/monomer-cosmos-sdk/server/cmd"

	"github.com/joshklop/gm-monomer/app"
	"github.com/joshklop/gm-monomer/cmd/gm-monomerd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
