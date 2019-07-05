package common

import (
	"encoding/json"
	"fmt"
	"log"
)

type definition struct {
	ID uint `json:"id"`
}

type postPayload struct {
	Definition   definition `json:"definition"`
	Parameters   string     `json:"parameters,omitempty"`
	SourceBranch string     `json:"sourceBranch"`
}

// GetPostPayload constructs a POST payload body for a request to the Azure DevOps API, with details of which build ID
// to queue, from which branch, and any additional parameters.
func GetPostPayload(buildDefID uint, parameters map[string]string, branch string) (string, error) {
	payload := postPayload{
		Definition: definition{
			ID: buildDefID,
		},
		SourceBranch: branch,
	}

	if len(parameters) > 0 {
		payload.Parameters = "{"
		firstParameter := true

		for key, value := range parameters {
			if firstParameter {
				firstParameter = false
			} else {
				payload.Parameters += ", "
			}
			payload.Parameters += fmt.Sprintf(`"Parameters.%s":"%s"`, key, value)
		}

		payload.Parameters += "}"
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
