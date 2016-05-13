package base

import(
  "net/http"
)

type Response struct {
  httpResp *http.Response
}

func (this *Response) HttpResp() *http.Response{
  return this.httpResp
}
