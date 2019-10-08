package common

import (
	"context"

	"github.com/spf13/viper"
)

// mustGetStorageAccountKey retrieves an access key to the Azure Storage Account containing the Terraform state files.
func mustGetStorageAccountKey() string {
	sc, errGsc := getStorageClient(viper.GetString("azure.state.subscription"))
	if errGsc != nil {
		panic(errGsc)
	}

	alkr, errLk := sc.ListKeys(
		context.TODO(),
		viper.GetString("azure.state.resourceGroup"),
		viper.GetString("azure.state.storageAccount"),
	)
	if errLk != nil {
		panic(errLk)
	}

	return *(*alkr.Keys)[0].Value
}
