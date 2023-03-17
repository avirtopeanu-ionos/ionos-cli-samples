package compute

import (
	"github.com/spf13/cobra"
	"ionos-cli-samples/styles/cobra/cmd/compute/server"
)

func init() {
	Compute.AddCommand(server.Server)

	Compute.PersistentFlags().Int32P("depth", "d", 1, "Request depth")
}

var Compute = &cobra.Command{
	Use:   "compute",
	Short: "IONOS Cloud V6 API",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
