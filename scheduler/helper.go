package scheduler

import(
  "net/http"
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/analyzer"
  "github.com/zhchsf/crawler/downloader"
  "time"
)

type genHttpClient func() *http.Client
type genRequest func() base.Request

func NewDownloaderPool(total uint32, httpClientGenerator genHttpClient) downloader.DownloaderPool {
  pool, err := downloader.NewDownloaderPool(
    total,
    func() downloader.Downloader {
      var gen genHttpClient
      if httpClientGenerator != nil {
        gen = httpClientGenerator
      }else{
        gen = generatehttpClient
      }
      return downloader.NewDownloader(gen())
    },
  )
  if err != nil {
    panic(err)
  }
  return pool
}

func NewAnalyzerPool(total uint32) analyzer.AnalyzerPool {
  pool, err := analyzer.NewAnalyzerPool(total)
  if err != nil {
    panic(err)
  }
  return pool
}

func generatehttpClient() *http.Client {
  return &http.Client{
    Timeout: 2 * time.Second,
  }
}
