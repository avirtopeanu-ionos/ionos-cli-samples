package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ionos-cli-samples/styles/cobra/cmd/compute"
	"os"
)

var (
	Root = &cobra.Command{
		Use:   "ionosctl",
		Short: "ionosctl powered by cobra",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	cfgFile string
	//Version string
)

func Execute() {
	if err := Root.Execute(); err != nil {
		//must.Die(err.Error()) // Apparently if err != nil; the error is printed before entering this block
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ionosctl.yaml)")

	Root.AddCommand(compute.Compute)
	Root.SilenceUsage = true
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
		//log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
