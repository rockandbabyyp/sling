package sling

import (
	"encoding/json"
	"net/http"
)

// ResponseDecoder decodes http responses into struct values.
// Decode decodes the response into the value pointed to by v.
type ResponseDecoder interface {
	Decode(resp *http.Response, v interface{}) error
}

// jsonDecoder decodes http response JSON into a JSON-tagged struct value.
type jsonDecoder struct {
}

// Decode the Response Body into the value pointed to by v.
// Caller must provide a non-nil v and close the resp.Body.
func (d jsonDecoder) Decode(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
