package downloader

import(
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/tool/id"
  "net/http"
)

var idGenerater id_generator.IdGenerator = id_generator.NewIdGenerator()

type Downloader interface {
  Id() uint32
  Download(request base.Request) (*base.Response, error)
}

type myDownloader struct {
  id uint32
  httpClient *http.Client
}

func NewDownloader(httpClient *http.Client) Downloader {
  id := idGenerater.NextId()
  return &myDownloader{id: id, httpClient: httpClient}
}

func (this *myDownloader) Id() uint32 {
  return this.id
}

func (this *myDownloader) Download(request base.Request) (*base.Response, error) {
  resp, err := this.httpClient.Do(request.HttpReq())
  if err != nil {
    return nil, err
  }
  return base.NewResponse(resp), nil
}
