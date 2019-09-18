package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/spf13/viper"
)

//nolint
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
		panic(errStackPath)
	}

	xStack := strings.Split(stackPath, string(os.PathSeparator))
	if len(xStack) < 2 {
		log.Fatalf("stack path '%s' should be least 2 levels deep, below '%s'",
			stackPath, viper.GetString("stackPrefix"))
	}

	const configSubs = "azure.subscriptions"
	subs := viper.GetStringMapString(configSubs)
	stackSub := xStack[0]
	if _, found := subs[stackSub]; !found {
		log.Fatalf("the subscription key '%s' is not present under '%s' in your config", stackSub, configSubs)
	}

	stackParent := xStack[len(xStack)-2]
	stack := xStack[len(xStack)-1]
	stateKey := fmt.Sprintf("%s.%s.%s", viper.GetString("azure.state.keyPrefix"), stackParent, stack)

	// Switching subscriptions
	cmdSetAccountSA := exec.Command("az", "account", "set",
		fmt.Sprintf("--subscription=%s", viper.GetString("azure.state.subscription")))
	fmt.Printf("Switching subscriptions... ")
	errSetAccountSA := cmdSetAccountSA.Run()
	if errSetAccountSA != nil {
		log.Fatalf("'az' errored when setting current subscription to %s: %s",
			viper.GetString("azure.state.subscription"), errSetAccountSA)
	}
	fmt.Println("done.")

	// Get access key; enables programmatic access to the storage account
	cmdSAKeys := exec.Command("az", "storage", "account", "keys", "list", "--account-name",
		viper.GetString("azure.state.storageAccount"))
	var out bytes.Buffer
	cmdSAKeys.Stdout = &out
	fmt.Printf("Retrieving storage account key... ")
	errSAKeys := cmdSAKeys.Run()
	if errSAKeys != nil {
		log.Fatalf("'az' errored when fetching storage account keys: %s", errSAKeys)
	}
	fmt.Println("done.")
	outBytes := out.Bytes()

	// Parse key out of JSON
	var saKeys []struct {
		Value string
	}

	errUmKeys := json.Unmarshal(outBytes, &saKeys)
	if errUmKeys != nil {
		log.Fatalf("unmarshaling '%s': %s", string(outBytes), errUmKeys)
	}

	// Switch subscriptions to given target
	cmdSetAccountTarget := exec.Command("az", "account", "set", fmt.Sprintf("--subscription=%s", subs[stackSub]))
	fmt.Printf("Switching subscriptions again... ")
	errSetAccountTarget := cmdSetAccountTarget.Run()
	if errSetAccountTarget != nil {
		log.Fatalf("'az' errored when setting current subscription to %s: %s",
			subs[stackSub], errSetAccountTarget)
	}
	fmt.Println("done.")

	// Announce init
	fmt.Println("Initialising Terraform with following dynamic values:")
	// fmt.Printf("\taccess_key:\t\t%s\n", saKeys[0].Value) // Don't output secrets
	fmt.Printf("\tcontainer_name:\t\t%s\n", subs[stackSub]) // Container name matches target sub GUID
	fmt.Printf("\tkey:\t\t\t%s\n", stateKey)
	fmt.Printf("\tstorage_account:\t%s\n", viper.GetString("azure.state.storageAccount"))

	// Run the initialisation
	cmdInit := exec.Command("terraform", "init",
		fmt.Sprintf("--backend-config=access_key=%s", saKeys[0].Value),
		fmt.Sprintf("--backend-config=container_name=%s", subs[stackSub]),
		fmt.Sprintf("--backend-config=key=%s", stateKey),
		fmt.Sprintf("--backend-config=storage_account_name=%s", viper.GetString("azure.state.storageAccount")))

	var wg sync.WaitGroup

	stdout, errOut := cmdInit.StdoutPipe()
	if errOut != nil {
		panic(errOut)
	}
	stderr, errErr := cmdInit.StderrPipe()
	if errErr != nil {
		panic(errErr)
	}
	errStart := cmdInit.Start()
	if errStart != nil {
		panic(errStart)
	}

	chPrint := make(chan string)

	scanOut := bufio.NewScanner(stdout)
	wg.Add(1)
	go func() {
		for scanOut.Scan() {
			chPrint <- scanOut.Text()
		}
		wg.Done()
	}()

	scanErr := bufio.NewScanner(stderr)
	wg.Add(1)
	go func() {
		for scanErr.Scan() {
			chPrint <- scanErr.Text()
		}
		wg.Done()
	}()

	var exitStatus int

	go func() {
		errWait := cmdInit.Wait()
		if errWait != nil {
			if exitErr, ok := errWait.(*exec.ExitError); ok {
				if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
					exitStatus = status.ExitStatus()
					log.Printf("Exit status: %d", exitStatus)
				}
			} else {
				panic(errWait)
			}
		}
		wg.Wait()
		close(chPrint)
	}()

	for line := range chPrint {
		fmt.Println(line)
	}

	os.Exit(exitStatus)
}
