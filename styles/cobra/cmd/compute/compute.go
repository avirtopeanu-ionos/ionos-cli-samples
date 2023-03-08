package compute

import (
	"github.com/spf13/cobra"
)

func init() {
	Compute.AddCommand(server)
}

var Compute = &cobra.Command{
	Use:   "compute",
	Short: "IONOS Cloud V6 API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
