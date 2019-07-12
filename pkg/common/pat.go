package common

import (
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

// ValidateAzureDevOpsPAT makes sure that the given token matches the format of an Azure DevOps Personal Access Token.
func ValidateAzureDevOpsPAT(token string) (string, error) {
	viperKey := "azureDevOps.pat"
	viperValue := viper.GetString(viperKey)
	if len(viperValue) == 0 {
		return "", fmt.Errorf(
			"the value for the '%s' key in the '%s' file has a length of zero",
			viperKey,
			viper.ConfigFileUsed())
	}

	rx := regexp.MustCompile("^[a-z0-9]{52}$")
	if !rx.Match([]byte(token)) {
		return "", fmt.Errorf(
			"the '%s' key in the '%s' file is not correctly formed",
			viperKey,
			viper.ConfigFileUsed())
	}

	return token, nil
}
