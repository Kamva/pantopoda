package pantopoda

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Pantopoda is a HTTP client that makes it easy to send HTTP requests and
// trivial to integrate with web services.
type Pantopoda struct {
}

// NewPantopoda generate new instance of pantopoda client
func NewPantopoda() *Pantopoda {
	return &Pantopoda{}
}

// Request sends a `method` request to the `endpoint` with given request data.
func (c *Pantopoda) Request(method string, endpoint string, request Request) (Response, error) {
	b := request.Payload.ToJSON()

	if !request.Query.Empty() {
		endpoint = endpoint + "?" + request.Query.ToString()
	}
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return Response(resBody), nil
}

// Get sends a GET request to `endpoint` with given data.
func (c *Pantopoda) Get(endpoint string, request Request) (Response, error) {
	return c.Request("GET", endpoint, request)
}

// Post sends a POST request to `endpoint` with given data.
func (c *Pantopoda) Post(endpoint string, request Request) (Response, error) {
	return c.Request("POST", endpoint, request)
}

// Put sends a PUT request to `endpoint` with given given data.
func (c *Pantopoda) Put(endpoint string, request Request) (Response, error) {
	return c.Request("PUT", endpoint, request)
}

// Patch sends a PATCH request to `endpoint` with given given data.
func (c *Pantopoda) Patch(endpoint string, request Request) (Response, error) {
	return c.Request("PATCH", endpoint, request)
}

// Delete sends a DELETE request to `endpoint` with given given data.
func (c *Pantopoda) Delete(endpoint string, request Request) (Response, error) {
	return c.Request("DELETE", endpoint, request)
}
