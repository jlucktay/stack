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

	payload, errPayload := common.GetPostPayload(123, map[string]string{"key": "value"}, "branch")
	is.NoErr(errPayload)

	req, errCreate := common.CreateBuildRequest(
		"dev.azure.com/build/me/a/build",
		"this is my PAT; there are many like it, but this one is mine",
		payload)
	is.NoErr(errCreate)

	handler := func(_ http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		is.True(strings.HasPrefix(auth, "Basic")) // 'Authorization' header value starts with 'Basic'.

		ct := r.Header.Get("Content-Type")
		is.Equal("application/json", ct) // 'Content-Type' == 'application/json'.
	}

	w := httptest.NewRecorder()
	handler(w, req)
}
