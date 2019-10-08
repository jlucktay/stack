package common

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	// AuthFromEnvClient is an env variable supported by the Azure SDK
	AuthFromEnvClient = "AZURE_CLIENT_ID"

	// AuthFromEnvTenant is an env variable supported by the Azure SDK
	AuthFromEnvTenant = "AZURE_TENANT_ID"

	// AuthFromFile is an env variable supported by the Azure SDK
	AuthFromFile = "AZURE_AUTH_LOCATION"
)

// newAuthorizer creates an Azure authorizer adhering to standard auth mechanisms provided by the Azure Go SDK.
// See Azure Go Auth docs here: https://docs.microsoft.com/en-us/go/azure/azure-sdk-go-authorization
func newAuthorizer() (*autorest.Authorizer, error) {
	// Carry out env var lookups
	_, clientIDExists := os.LookupEnv(AuthFromEnvClient)
	_, tenantIDExists := os.LookupEnv(AuthFromEnvTenant)
	_, fileAuthSet := os.LookupEnv(AuthFromFile)

	// Execute logic to return an authorizer from the correct method
	switch {
	case clientIDExists && tenantIDExists:
		authorizer, err := auth.NewAuthorizerFromEnvironment()
		return &authorizer, err
	case fileAuthSet:
		authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
		return &authorizer, err
	default:
		authorizer, err := auth.NewAuthorizerFromCLI()
		return &authorizer, err
	}
}

// getStorageClient is a helper function that will setup an Azure Storage Account client on your behalf.
func getStorageClient(subscriptionID string) (*storage.AccountsClient, error) {
	// Create a Storage Account client
	storageAccountClient := storage.NewAccountsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := newAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	storageAccountClient.Authorizer = *authorizer

	return &storageAccountClient, nil
}
