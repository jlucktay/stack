package common_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jlucktay/stack/pkg/common"
)

func TestBuildRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello %s\n", "world")

		for key, value := range r.Header {
			fmt.Fprintf(w, "'%s'='%#v'\n", key, value)
		}
	}

	payload, errPayload := common.GetPostPayload(123, map[string]string{"key": "value"}, "branch")
	if errPayload != nil {
		t.Fatal(errPayload)
	}
	req, errCreate := common.CreateBuildRequest(
		"dev.azure.com/build/me/a/build",
		"this is my PAT, there are many like it, but this one is mine",
		payload)
	if errCreate != nil {
		t.Fatal(errCreate)
	}
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	t.Log(resp.StatusCode)
	t.Logf("%#v\n", resp.Header)
	t.Log(string(body))
}
