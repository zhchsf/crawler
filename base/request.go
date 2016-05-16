package base

import(
  "net/http"
  "strings"
)

type Request struct{
  url string

  // html json text
  // now not usable
  respType string

  // get post
  method string

  // post data "a=1&b=2", url.Values Encode()
  postData string

  header http.Header

  cookies []*http.Cookie

  httpReq *http.Request

  // proxy string
}

func NewRequest(url, respType, method, postData string, header http.Header, cookies []*http.Cookie) *Request {
  return &Request{url: url, respType: respType, method: method, header: header, cookies: cookies}
}

// func NewGetRequest()

// func (this *Request) AddProxy(proxy string) *Request {
//   this.proxy = proxy
//   return this
// }

func (this *Request) HttpReq() *http.Request {
  if this.httpReq != nil {
    return this.httpReq
  }

  method := strings.ToUpper(this.method)
  body := strings.NewReader(this.postData)
  req, err := http.NewRequest(method, this.url, body)
  if err != nil {
    panic(err)
  }
  if this.header != nil {
    req.Header = this.header
  }
  if this.cookies != nil {
    for _, cookie := range this.cookies {
      req.AddCookie(cookie)
    }
  }
  this.httpReq = req
  return req
}

func (this *Request) Valid() bool {
  return this.url != "" && this.method != ""
}
