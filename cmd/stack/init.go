package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/spf13/viper"
)

func initStack() {
	stackPath, errStackPath := common.GetStackPath(
		viper.GetString("stackPrefix"),
		fmt.Sprintf(
			"github.com/%s/%s",
			viper.GetString("github.org"),
			viper.GetString("github.repo"),
		),
	)
	if errStackPath != nil {
		log.Fatal(errStackPath)
	}

	fmt.Printf("stackPath: %#v\n", stackPath)

	xStack := strings.Split(stackPath, string(os.PathSeparator))
	if len(xStack) < 3 {
		log.Fatalf("stack path '%s' should have be least 3 levels deep", stackPath)
	}

	const configSubs = "azure.subscriptions"
	subs := viper.GetStringMapString(configSubs)
	stackSub := xStack[0]
	if _, found := subs[stackSub]; !found {
		log.Fatalf("the subscription key '%s' is not present under '%s' in your config", stackSub, configSubs)
	}

	fmt.Printf("stackSub: %#v\n", stackSub)
	fmt.Printf("subscription GUID: %#v\n", subs[stackSub])

	stackParent := xStack[len(xStack)-2]
	stack := xStack[len(xStack)-1]
	stateKey := fmt.Sprintf("%s.%s.%s", viper.GetString("azure.stateKeyPrefix"), stackParent, stack)

	fmt.Printf("stateKey: '%s'\n", stateKey)
	fmt.Printf("storage account: '%s'\n", viper.GetString("azure.stateStorageAccount"))
}
