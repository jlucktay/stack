package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type buildLinkWeb struct {
	Href string `json:"href"`
}

type buildLink struct {
	Web buildLinkWeb `json:"web"`
}

type build struct {
	Links buildLink `json:"_links"`
}

// SendBuildRequest will POST a HTTP request with the given payload body to the given Azure DevOps API URL, using the
// given token for authentication.
// A successfully queued build will return its URL.
func SendBuildRequest(apiurl, token, reqBody string) (*url.URL, error) {
	req, errReq := http.NewRequest(http.MethodPost, apiurl, strings.NewReader(reqBody))
	if errReq != nil {
		return nil, errReq
	}

	req.SetBasicAuth("username", token) // Azure DevOps API doesn't care about the username, just the token
	req.Header.Add("Content-Type", "application/json")

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("URL '%s': status code error: %d %s\n"+
			"Does your Azure DevOps token have the 'Build (Read & execute)' scope?",
			apiurl, resp.StatusCode, resp.Status)
	}

	respBytes, errByteRead := ioutil.ReadAll(resp.Body)
	if errByteRead != nil {
		return nil, errByteRead
	}

	var buildResult build

	errUm := json.Unmarshal(respBytes, &buildResult)
	if errUm != nil {
		return nil, errUm
	}

	buildURL, errParse := url.Parse(buildResult.Links.Web.Href)
	if errParse != nil {
		return nil, errParse
	}

	return buildURL, nil
}
