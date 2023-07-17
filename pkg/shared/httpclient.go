package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-resty/resty/v2"
)

// HTTPClient is an interface that allows us to implements any
// http client library we want
type HTTPClient interface {
	Do(req *http.Request, resp any) (*http.Response, error)
}

// httpClient is a resty implementation of HTTPClient interface
type httpClient struct {
	restyHttp *resty.Client
}

// NewHTTPClient returns a new instance of httpClient
func NewHTTPClient() *httpClient {
	client := resty.New()
	client.OnAfterResponse(LogAfterResponse)
	client.OnBeforeRequest(LogBeforeRequest)
	return &httpClient{client}
}

// Do execute the request and returns response and error
//
// *http.Response: is a resty rawResponse, so it has a closed body to avoid memory leaks
// for that reason this method uses response any for body mapping letting to resty the
// responsibility to map and close the body
//
// error: is a resty error but always *http.Response is sent to the caller to handle http.Status
func (c *httpClient) Do(request *http.Request, response any) (*http.Response, error) {
	headers := make(map[string]string)
	for k, v := range request.Header {
		headers[k] = v[0]
	}

	res, err := c.restyHttp.R().
		SetBody(request.Body).
		SetHeaders(headers).
		SetResult(response). // TODO: improve to avoid nil pointer exception
		Execute(request.Method, request.URL.String())
	if err != nil {
		return nil, err
	}

	return res.RawResponse, nil
}

type ContextKey string

const (
	LogHttpClient            = "shared/httpclient.go"
	RequestUUID   ContextKey = "request_uuid"
)

// LogAfterResponse is a resty callback that logs the response of the request
// it uses the request context to get the requestID to sign the log and make easy
// to track the request in the logs
func LogAfterResponse(r *resty.Client, response *resty.Response) error {
	requestID := fmt.Sprintf("%v", response.Request.Context().Value(RequestUUID))
	LogResponse(requestID, response.Status(), string(response.Body()), response.Request.Method, response.Request.URL, response.Header())
	return nil
}

// LogBeforeRequest is a resty callback that logs the request before it is sent
// it uses the request context to set the requestID to sign the log and make easy
// to track the request in the logs
func LogBeforeRequest(_ *resty.Client, request *resty.Request) error {
	id, err := uuid.NewUUID()
	if err != nil {
		LogError("error uuid lib fail", LogHttpClient, "LogBeforeRequest", err)
	}

	ctx := context.WithValue(request.Context(), RequestUUID, id.String())
	request.SetContext(ctx)

	body, err := json.Marshal(request.Body)
	if err != nil {
		LogError("error json marshal fail", LogHttpClient, "LogBeforeRequest", err)
	}
	LogRequest(id.String(), request.Method, request.URL, string(body), request, request.Header)
	return nil
}
