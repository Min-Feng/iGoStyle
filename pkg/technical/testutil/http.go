package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"

	"AmazingTalker/pkg/technical/types"
)

func HTTPResponseBody(router http.Handler, httpMethod string, url string, body io.Reader) string {
	wRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(httpMethod, url, body)
	router.ServeHTTP(wRecorder, req)
	actualBody := types.StringTool{}.ToPrettyJSON(wRecorder.Body.Bytes())
	return actualBody
}
