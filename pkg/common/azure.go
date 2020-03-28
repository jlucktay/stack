package common

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/spf13/viper"
)

// mustGetStorageAccountKey retrieves an access key to the Azure Storage Account containing the Terraform state files.
func mustGetStorageAccountKey() string {
	subKey := "azure.state.subscription"
	if !viper.IsSet(subKey) {
		panic("the GUID of the subscription containing the remote state storage account " +
			"has not been specified under '" + subKey + "' in your config")
	}

	sc, errGsc := getStorageClient(viper.GetString(subKey))
	if errGsc != nil {
		panic(errGsc)
	}

	rgKey := "azure.state.resourceGroup"
	if !viper.IsSet(rgKey) {
		panic("the name of the resource group containing the remote state storage account " +
			"has not been specified under '" + rgKey + "' in your config")
	}

	saKey := "azure.state.storageAccount"
	if !viper.IsSet(saKey) {
		panic("the name of the remote state storage account has not been specified under '" +
			saKey + "' in your config")
	}

	alkr, errLk := sc.ListKeys(context.TODO(), viper.GetString(rgKey), viper.GetString(saKey), storage.Kerb)
	if errLk != nil {
		panic(errLk)
	}

	return *(*alkr.Keys)[0].Value
}
