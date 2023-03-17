package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"ionos-cli-samples/pkg/client"
	"ionos-cli-samples/pkg/must"
	"ionos-cli-samples/styles/cobra/internal/constants"
	"ionos-cli-samples/styles/cobra/internal/printer"
)

func init() {
}

var ls = &cobra.Command{
	Use:   "list",
	Short: "list servers",
	PreRunE: func(c *cobra.Command, args []string) error {
		return errors.Join(
			c.MarkFlagRequired(constants.FlagIdDatacenter),
		)
	},
	RunE: func(c *cobra.Command, args []string) error {
		dcs, resp, err := client.Compute().ServersApi.DatacentersServersGet(
			context.Background(),
			must.Get(c.Flags().GetString(constants.FlagIdDatacenter)),
		).Depth(must.Get(c.Flags().GetInt32("depth"))).Execute()
		if err != nil {
			return fmt.Errorf("failed getting servers: %s", resp.Status)
		}

		//json, err := json.Marshal(dcs)
		//if err != nil {
		//	return err
		//}

		//fmt.Printf(string(json) + "\n")
		printer.PrintTable((*dcs.Items)[0])
		return nil
	},
}
