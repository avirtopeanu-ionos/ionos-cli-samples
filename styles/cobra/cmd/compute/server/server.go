package server

import (
	"github.com/spf13/cobra"
	"ionos-cli-samples/styles/cobra/internal/constants"
)

func init() {
	Server.PersistentFlags().StringP(constants.FlagIdDatacenter, constants.FlagIdShort, "", "")

	Server.AddCommand(ls)
}

var Server = &cobra.Command{
	Use:   "server",
	Short: "manage compute servers",
	Run: func(c *cobra.Command, args []string) {
	},
}
