package application

import (
	"io"
	"net/http"
	"encoding/json"
)

type Request struct {
	source *http.Request
}

func NewRequest(source *http.Request) *Request {
	return &Request{source: source}
}

func (r *Request) Data() io.Reader {
	return r.source.Body
}

type Response struct {
	HeaderKey string
	HeaderValue string
	Result []byte
}

func NewResponse(result interface{}, err error) (*Response, error) {
	encoded, converted := json.Marshal(result)
	w := &Response{}
	w.Result = encoded
	w.HeaderKey = "Content-Type"
	w.HeaderValue = "application/json"
	if err == nil {err = converted}
	return w, err
}
