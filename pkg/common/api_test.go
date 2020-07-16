package common_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/stack/pkg/common"
)

func TestBuildRequest(t *testing.T) {
	is := is.New(t)

	handler := func(_ http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		is.True(strings.HasPrefix(auth, "Basic"))                  // 'Authorization' header value starts with 'Basic'.
		is.Equal("application/json", r.Header.Get("Content-Type")) // 'Content-Type' == 'application/json'.
	}

	var testBuildID uint = 123

	payload, errPayload := common.GetPostPayload(testBuildID, map[string]string{"key": "value"}, "branch")
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
