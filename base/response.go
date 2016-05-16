package base

import(
  "net/http"
)

type Response struct {
  httpResp *http.Response
}

func NewResponse(resp *http.Response) *Response {
  return &Response{httpResp: resp}
}

func (this *Response) HttpResp() *http.Response{
  return this.httpResp
}

func (this *Response) Valid() bool {
  return this.httpResp != nil
}
