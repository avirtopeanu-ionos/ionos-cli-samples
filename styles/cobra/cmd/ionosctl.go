package cmd

import (
	"fmt"
	sdk "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ionos-cli-samples/pkg/die"
	"log"
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
	Client  *sdk.APIClient
	//Version string
)

func Execute() {
	if err := Root.Execute(); err != nil {
		die.Die(err.Error())
	}
}
func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.Print("Reading config...")
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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	Client = sdk.NewAPIClient(
		sdk.NewConfiguration(
			viper.GetString("IONOS_USERNAME"),
			viper.GetString("IONOS_PASSWORD"),
			viper.GetString("IONOS_TOKEN"),
			viper.GetString("IONOS_API_URL"),
		),
	)

	log.Println(Client.GetConfig().Token)
}
