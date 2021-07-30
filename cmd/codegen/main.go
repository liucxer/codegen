package main

import (
	"fmt"
	"github.com/liucxer/codegen/cmd/codegen/gen"
	"github.com/liucxer/codegen/cmd/codegen/hook"
	"github.com/liucxer/codegen/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var verbose = false

var cmdRoot = &cobra.Command{
	Use:     "codegen",
	Version: version.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	},
}

func init() {
	cmdRoot.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "")

	cmdRoot.AddCommand(gen.CmdGen)
	cmdRoot.AddCommand(hook.CmdHook)
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
