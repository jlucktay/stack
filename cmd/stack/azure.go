package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

// mustGetStorageAccountKey uses the Azure CLI to retrieve an access key for the Azure Storage Account containing the
// Terraform state files.
func mustGetStorageAccountKey() string {
	cmdSAKeys := exec.Command("az", "storage", "account", "keys", "list", "--account-name",
		viper.GetString("azure.state.storageAccount"))
	var out bytes.Buffer
	cmdSAKeys.Stdout = &out
	fmt.Printf("Retrieving storage account key... ")
	errSAKeys := cmdSAKeys.Run()
	if errSAKeys != nil {
		panic(fmt.Sprintf("'az' errored when fetching storage account keys: %s", errSAKeys))
	}
	fmt.Println("done.")
	outBytes := out.Bytes()

	// Parse key out of JSON
	var saKeys []struct {
		Value string
	}

	errUmKeys := json.Unmarshal(outBytes, &saKeys)
	if errUmKeys != nil {
		panic(fmt.Sprintf("unmarshaling '%s': %s", string(outBytes), errUmKeys))
	}

	return saKeys[0].Value
}
