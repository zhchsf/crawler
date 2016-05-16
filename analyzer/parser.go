package analyzer

import(
  "net/http"
  "github.com/zhchsf/crawler/base"
)

type ParseResponse func(resp *http.Response) ([]base.Data, []error)
