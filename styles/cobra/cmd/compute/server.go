package compute

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	Compute.AddCommand(server)
}

var server = &cobra.Command{
	Use:   "server",
	Short: "manage compute servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Hello!\n")
		return nil
	},
}
