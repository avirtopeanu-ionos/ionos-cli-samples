package compute

import (
	"github.com/spf13/cobra"
	"ionos-cli-samples/styles/cobra/cmd"
)

func init() {
	cmd.Root.AddCommand(Compute)
}

var Compute = &cobra.Command{
	Use:   "compute",
	Short: "IONOS Cloud V6 API",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
