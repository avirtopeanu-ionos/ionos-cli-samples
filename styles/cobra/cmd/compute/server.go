package compute

import (
	"context"
	"github.com/spf13/cobra"
	"ionos-cli-samples/pkg/client"
	"ionos-cli-samples/pkg/must"
	"ionos-cli-samples/styles/cobra/internal/constants"
)

func init() {
	server.Flags().StringP(constants.FlagIdDatacenter, constants.FlagIdShort, "", "")
}

var server = &cobra.Command{
	Use:   "server",
	Short: "manage compute servers",
	RunE: func(c *cobra.Command, args []string) error {
		_, _, _ = client.Compute().ServersApi.DatacentersServersGet(context.Background(), must.Get(c.Flags().GetString(constants.FlagIdDatacenter))).Execute()
		return nil
	},
}
