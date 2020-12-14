package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"

	"iGoStyle/pkg/technical/types"
)

func HTTPResponse(router http.Handler, httpMethod string, url string, body io.Reader) string {
	wRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(httpMethod, url, body)
	router.ServeHTTP(wRecorder, req)
	actualBody := types.StringUtil{}.ToPrettyJSON(wRecorder.Body.Bytes())
	return actualBody
}
