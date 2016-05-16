package analyzer

import(
  "github.com/zhchsf/crawler/tool/pool"
  "reflect"
  "errors"
)

type AnalyzerPool interface {
  Take() (Analyzer, error)
  Return(Analyzer) error
  Total() uint32
  Used() uint32
}

type myAnalyzerPool struct {
  pool pool.Pool
  etype reflect.Type
}

func NewAnalyzerPool(total uint32) (AnalyzerPool, error) {
  gen := NewAnalyzer
  etype := reflect.TypeOf(gen())
  genEntity := func() pool.Entity {
    return gen()
  }
  pool, err := pool.NewPool(total, etype, genEntity)
  if err != nil {
    return nil, err
  }
  return &myAnalyzerPool{pool: pool, etype: etype}, nil
}

func (this *myAnalyzerPool) Take() (Analyzer, error) {
  entity, err := this.pool.Take()
  if err != nil {
    return nil, err
  }
  analyzer, ok := entity.(Analyzer)
  if !ok {
    panic(errors.New("trans error"))
  }
  return analyzer, nil
}

func (this *myAnalyzerPool) Return(analyzer Analyzer) error {
  return this.pool.Return(analyzer)
}

func (this *myAnalyzerPool) Total() uint32{
  return this.pool.Total()
}

func (this *myAnalyzerPool) Used() uint32{
  return this.pool.Used()
}

