package scheduler

import(
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/analyzer"
  "github.com/zhchsf/crawler/downloader"
  "time"
  "fmt"
)

type Scheduler struct {
  analyzerPool analyzer.AnalyzerPool
  downloaderPool downloader.DownloaderPool
  respChan chan base.Response
  reqChan chan base.Request
  timeInterval time.Duration // schedule时间间隔，可以设置为毫秒级
  respParsers []analyzer.ParseResponse
  requestGenerator genRequest // 生产request func，返回base.Request，外部逻辑自己实现并传入
}

func NewScheduler() *Scheduler {
  return &Scheduler{}
}

func (this *Scheduler) Start(
  poolConfig base.PoolConfig, 
  chanConfig base.ChanConfig, 
  httpClientGenerator genHttpClient,
  respParsers []analyzer.ParseResponse,
  timeInterval time.Duration,
  requestGenerator genRequest) {
  this.downloaderPool = NewDownloaderPool(poolConfig.DownloaderTotal, httpClientGenerator)
  this.analyzerPool = NewAnalyzerPool(poolConfig.AnalyzerTotal)
  this.respChan = make(chan base.Response, chanConfig.RespChanLen)
  this.reqChan = make(chan base.Request, chanConfig.ReqChanLen)
  this.timeInterval = timeInterval
  this.respParsers = respParsers
  this.requestGenerator = requestGenerator

  this.startDownloading()
  this.startAnalyzing()
  this.schedule()
}

func (this *Scheduler) startDownloading() {
  go func(){
    for{
      req, ok := <-this.reqChan
      if !ok {
        break
      }
      go this.download(req)
    }
  }()
}

func (this *Scheduler) download(req base.Request) {
  downloader, err := this.downloaderPool.Take()
  if err != nil {
    //
  }
}

func (this *Scheduler) startAnalyzing() {
  go func(){
    for{
      resp, ok := <-this.respChan
      if !ok {
        break
      }
      go this.analyze(resp)
    }
  }()
}

func (this *Scheduler) analyze(resp base.Response) {
  fmt.Println(resp.HttpResp().Body)
  defer resp.HttpResp().Body.Close()
}

func (this *Scheduler) schedule() {
  for{
    remainder := cap(this.reqChan) - len(this.reqChan)
    for remainder > 0 {
      this.reqChan <- this.requestGenerator()
      remainder--
    }
    time.Sleep(this.timeInterval)
  }
}
