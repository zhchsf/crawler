package base

import(
  "net/http"
)

type Request struct{
  url string

  // html json text
  respType string

  // get post
  method string

  // post data
  postData string

  header http.Header

  cookies []*http.Cookie

  proxy string
}

func NewRequest(url, respType, method, postData string, header http.Header, cookies []*http.Cookie) *Request {
  return &Request{url: url, respType: respType, method: method, header: header, cookies: cookies}
}

func (this *Request) AddProxy(proxy string) *Request {
  this.proxy = proxy
  return this
}
