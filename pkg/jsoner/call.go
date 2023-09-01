package jsoner

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

func Call(method, uri, body string, headers map[string]string) (int, map[string]interface{}, error) {
	return CallContext(context.Background(), method, uri, body, headers)
}

func CallContext(ctx context.Context, method, uri, body string, headers map[string]string) (int, map[string]interface{}, error) {
	code := -1
	response := make(map[string]interface{})
	request, reqErr := http.NewRequestWithContext(ctx, method, uri, bytes.NewBuffer([]byte(body)))
	if reqErr != nil {
		return code, response, reqErr
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	result, resErr := new(http.Client).Do(request)
	if resErr != nil {
		return code, response, resErr
	}
	buffer, bufErr := ioutil.ReadAll(result.Body)
	if bufErr != nil {
		return code, response, bufErr
	}
	code = result.StatusCode
	if marshalErr := Unmarshal(buffer, &response); marshalErr != nil {
		return code, response, marshalErr
	}
	if closeErr := result.Body.Close(); closeErr != nil {
		return code, response, closeErr
	}
	return code, response, nil
}
