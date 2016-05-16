package scheduler

import(
  "net/http"
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/analyzer"
  "github.com/zhchsf/crawler/downloader"
)

type genHttpClient func() *http.Client
type genRequest func() base.Request

func NewDownloaderPool(total uint32, httpClientGenerator genHttpClient) downloader.DownloaderPool {
  pool, err := downloader.NewDownloaderPool(
    total,
    func() downloader.Downloader {
      return downloader.NewDownloader(httpClientGenerator())
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
