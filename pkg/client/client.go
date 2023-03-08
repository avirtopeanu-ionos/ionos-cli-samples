package client

import (
	compute "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Map

// Compute returns the initialized Compute API V6 client
func Compute() *compute.APIClient {
	client, ok := once.Load("compute")
	if ok {
		return client.(*compute.APIClient)
	}

	client = loadComputeClient()
	once.Store("compute", client)
	return client.(*compute.APIClient)
}

func loadComputeClient() *compute.APIClient {
	conf := compute.NewConfiguration(
		viper.GetString("IONOS_USERNAME"),
		viper.GetString("IONOS_PASSWORD"),
		viper.GetString("IONOS_TOKEN"),
		viper.GetString("IONOS_API_URL"),
	)
	conf.SetDepth(1)
	return compute.NewAPIClient(conf)
}

// GetPostgresClient
// loadPostgresClient
