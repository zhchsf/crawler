package analyzer

import(
  "github.com/zhchsf/crawler/base"
  "github.com/zhchsf/crawler/tool/id"
  "errors"
)

var idGenerater id_generator.IdGenerator = id_generator.NewIdGenerator()

type Analyzer interface {
  Id() uint32
  Analyze(respParsers []ParseResponse, response base.Response) ([]base.Data, []error)
}

type myAnalyzer struct {
  id uint32
}

func NewAnalyzer() Analyzer {
  return &myAnalyzer{id: idGenerater.NextId()}
}

func (this *myAnalyzer) Id() uint32 {
  return this.id
}

// Todo
func (this *myAnalyzer) Analyze(respParsers []ParseResponse, response base.Response) ([]base.Data, []error) {
  if respParsers == nil {
    return nil, []error{errors.New("nil parsers")}
  }
  resp := response.HttpResp()
  if resp == nil {
    return nil, []error{errors.New("nil response")}
  }

  dataList := make([]base.Data, 0)
  errors := make([]error, 0)
  for _, parser := range respParsers {
    datas, errs := parser(resp)
    for _, pData := range datas {
      dataList = append(dataList, pData)
    }
    for _, err := range errs {
      errors = append(errors, err)
    }
  }
  return dataList, errors
}
