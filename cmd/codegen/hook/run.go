package hook

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func init() {
	CmdHook.AddCommand(cmdHookRun)
}

var cmdHookRun = &cobra.Command{
	Use:   "run",
	Short: "git hook run",
	Run: func(cmd *cobra.Command, args []string) {
		spew.Dump(args)
		os.Exit(1)
	},
}
