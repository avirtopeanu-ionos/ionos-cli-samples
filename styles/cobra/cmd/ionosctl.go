package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ionos-cli-samples/pkg/must"
	"ionos-cli-samples/styles/cobra/cmd/compute"
	"os"
)

var (
	Root = &cobra.Command{
		Use:   "ionosctl",
		Short: "ionosctl powered by cobra",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cfgFile string
	//Version string
)

func Execute() {
	if err := Root.Execute(); err != nil {
		must.Die(err.Error())
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ionosctl.yaml)")

	Root.AddCommand(compute.Compute)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ionosctl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
