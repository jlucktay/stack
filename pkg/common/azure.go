package common

import (
	"context"

	"github.com/spf13/viper"
)

// mustGetStorageAccountKey retrieves an access key to the Azure Storage Account containing the Terraform state files.
func mustGetStorageAccountKey() string {
	sub := viper.GetString("azure.state.subscription")
	if sub == "" {
		panic("the GUID of the subscription containing the remote state storage account " +
			"has not been specified in your config file")
	}

	sc, errGsc := getStorageClient(sub)
	if errGsc != nil {
		panic(errGsc)
	}

	rg := viper.GetString("azure.state.resourceGroup")
	if rg == "" {
		panic("the name of the resource group containing the remote state storage account" +
			"has not been specified in your config file")
	}
	sa := viper.GetString("azure.state.storageAccount")
	if sa == "" {
		panic("the name of the remote state storage account has not been specified in your config file")
	}

	alkr, errLk := sc.ListKeys(context.TODO(), rg, sa)
	if errLk != nil {
		panic(errLk)
	}

	return *(*alkr.Keys)[0].Value
}
