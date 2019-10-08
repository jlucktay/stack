package common

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func InitStack() {
	stackPath := mustGetStackPath()

	xStack := strings.Split(stackPath, string(os.PathSeparator))
	if len(xStack) < 2 {
		panic(fmt.Sprintf("stack path '%s' should be least 2 levels deep, below '%s'",
			stackPath, viper.GetString("stackPrefix")))
	}

	const configSubs = "azure.subscriptions"
	subs := viper.GetStringMapString(configSubs)
	stackSub := xStack[0]
	if _, found := subs[stackSub]; !found {
		panic(fmt.Sprintf("the subscription key '%s' is not present under '%s' in your config", stackSub, configSubs))
	}

	stackParent := xStack[len(xStack)-2]
	stack := xStack[len(xStack)-1]
	stateKey := fmt.Sprintf("%s.%s.%s", viper.GetString("azure.state.keyPrefix"), stackParent, stack)

	// Get access key; enables programmatic access to the storage account
	saKey := mustGetStorageAccountKey()

	// Announce init
	fmt.Println("Initialising Terraform with following dynamic values:")
	fmt.Printf("\tcontainer_name:\t\t%s\n", subs[stackSub]) // Container name matches target sub GUID
	fmt.Printf("\tkey:\t\t\t%s\n", stateKey)
	fmt.Printf("\tstorage_account:\t%s\n", viper.GetString("azure.state.storageAccount"))

	// Run the initialisation
	cmdInit := exec.Command("terraform", "init",
		fmt.Sprintf("--backend-config=access_key=%s", saKey),
		fmt.Sprintf("--backend-config=container_name=%s", subs[stackSub]),
		fmt.Sprintf("--backend-config=key=%s", stateKey),
		fmt.Sprintf("--backend-config=storage_account_name=%s", viper.GetString("azure.state.storageAccount")))

	run(cmdInit)
}
