func ClustersNodepoolsDeleteCmd() *core.Command {
	cmd := core.NewCommand(context.TODO(), nil, core.CommandBuilder{
		Namespace: "dataplatform",
		Resource:  "nodepool",
		Verb:      "delete",
		Aliases:   []string{"del", "d"},
		ShortDesc: "Remove a DataPlatformNodePool from a DataPlatformCluster",
		Example:   "ionosctl dataplatform nodepool delete --cluster-id <cluster-id>",
		PreCmdRun: func(c *core.PreCommandConfig) error {
			return core.CheckRequiredFlagsSets(c.Command, c.NS, []string[]string{constants.ArgAll}, []string{constants.FlagClusterId}, []string{constants.ArgAll, constants.FlagClusterId})
		},
		CmdRun: func(c *core.CommandConfig) error {
			// Implement the actual command logic here
		},
		InitClient: true,
	})

	// TODO: Check me! Did I successfully add all flags for ClustersNodepoolsDelete?
	


	return cmd
}
