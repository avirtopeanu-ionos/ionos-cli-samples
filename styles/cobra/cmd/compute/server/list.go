package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"ionos-cli-samples/pkg/client"
	"ionos-cli-samples/pkg/must"
	"ionos-cli-samples/styles/cobra/internal/constants"
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
		_, resp, err := client.Compute().ServersApi.DatacentersServersGet(
			context.Background(),
			must.Get(c.Flags().GetString(constants.FlagIdDatacenter)),
		).Execute()
		if err != nil {
			return fmt.Errorf("failed getting servers: %s", resp.Status)
		}

		json, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		fmt.Printf(string(json) + "\n")
		return nil
	},
}
