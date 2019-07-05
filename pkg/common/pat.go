package common

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type personalAccessToken struct {
	Token string `json:"token"`
}

var patPath = os.ExpandEnv("${HOME}/.config/dan-clz-mig-fac/pat.json")

// getPAT looks in the user's XDG .config dir for a JSON file containing their Azure DevOps Personal Access Token
func getPAT() (string, error) {
	jsonFile, errOpen := os.Open(patPath)
	if errOpen != nil {
		fmt.Fprintf(os.Stderr, "Could not find '%s'!\n"+
			"Please create and populate the 'pat' key's value with a Personal Access Token from Azure DevOps:\n"+
			"- https://dev.azure.com/DanClzAutomation/_usersSettings/tokens\n"+
			"This token requires the 'Build (Read & execute)' scope only.", patPath)
		return "", fmt.Errorf("could not find '%s'", patPath)
	}
	token := personalAccessToken{}
	dec := json.NewDecoder(jsonFile)
	if errDecode := dec.Decode(&token); errDecode != nil {
		return "", errDecode
	}

	return validatePAT(token.Token)
}

// validatePAT makes sure that the given token matches the format of Azure DevOps
func validatePAT(token string) (string, error) {
	if len(token) == 0 {
		return "", fmt.Errorf("the token has a length of zero")
	}
	rx := regexp.MustCompile("^[a-z0-9]{52}$")
	if !rx.Match([]byte(token)) {
		return "", fmt.Errorf("the 'pat' key in the '%s' file is not correctly formed", patPath)
	}

	return token, nil
}
