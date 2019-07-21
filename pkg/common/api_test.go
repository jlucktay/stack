package common_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/matryer/is"
)

func TestBuildRequest(t *testing.T) {
	is := is.New(t)

	handler := func(_ http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		is.True(strings.HasPrefix(auth, "Basic"))                  // 'Authorization' header value starts with 'Basic'.
		is.Equal("application/json", r.Header.Get("Content-Type")) // 'Content-Type' == 'application/json'.
	}

	payload, errPayload := common.GetPostPayload(123, map[string]string{"key": "value"}, "branch")
	if errPayload != nil {
		t.Fatal(errPayload)
	}
	req, errCreate := common.CreateBuildRequest(
		"dev.azure.com/build/me/a/build",
		"this is my PAT; there are many like it, but this one is mine",
		payload)
	if errCreate != nil {
		t.Fatal(errCreate)
	}
	w := httptest.NewRecorder()
	handler(w, req)
}
