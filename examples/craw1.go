package main

import(
  "net/http"
  "time"
  "github.com/zhchsf/crawler/scheduler"
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/analyzer"
  "github.com/PuerkitoBio/goquery"
  "fmt"
  "io"
)

// func httpClientGenerator() *http.Client {
//   return &http.Client{
//     Timeout: time.Second,
//   }
// }

func requestGenerator() base.Request {
  url := "http://testing.caishuo.com" // /topics
  respType := "html"
  method := "get"
  postData := ""
  headers := map[string]string{
    "Referer": "http://www.baidu.com",
  }
  return *base.NewRequest(url, respType, method, postData, headers, nil)
}

// TODO
func parseForNews(httpResp *http.Response) ([]base.Data, []error) {
  var httpRespBody io.ReadCloser = httpResp.Body
  defer httpRespBody.Close()
  dataList := make([]base.Data, 0)
  errs := make([]error, 0)
  doc, _ := goquery.NewDocumentFromReader(httpRespBody)
  fmt.Println("-----------------------------------")
  doc.Find("#topics_list dt a em").Each(func(index int, sel *goquery.Selection){
    fmt.Println(sel.Text())
  })
  return dataList, errs
}

func getRespParsers() []analyzer.ParseResponse {
  parsers := []analyzer.ParseResponse{
    parseForNews,
  }
  return parsers
}

func main(){
  sched := scheduler.NewScheduler()
  // sched.SetHttpClient(httpClientGenerator)
  poolConfig := base.NewPoolConfig(4, 4)
  chanConfig := base.NewChanConfig(4, 4)
  respParsers := getRespParsers()
  timeInterval := 100 * time.Millisecond
  sched.Start(poolConfig, chanConfig, respParsers, timeInterval, requestGenerator)
}
