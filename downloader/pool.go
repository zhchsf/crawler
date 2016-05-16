package downloader

import(
  "github.com/zhchsf/crawler/tool/pool"
  "reflect"
  "errors"
)

type DownloaderPool interface {
  Take() (Downloader, error)
  Return(downloader Downloader) error
  Total() uint32
  Used() uint32
}

type myDownloaderPool struct {
  pool pool.Pool
  etype reflect.Type
}

func NewDownloaderPool(total uint32, gen func() Downloader) (DownloaderPool, error) {
  etype := reflect.TypeOf(gen())
  genEntity := func() pool.Entity {
    return gen()
  }
  pool, error := pool.NewPool(total, etype, genEntity)
  if error != nil {
    return nil, error
  }
  return &myDownloaderPool{pool: pool, etype: etype}, nil
}

func (this *myDownloaderPool) Take() (Downloader, error) {
  entity, error := this.pool.Take()
  if error != nil {
    return nil, error
  }
  downloader, ok := entity.(Downloader)
  if !ok {
    panic(errors.New("类型转换出错"))
  }
  return downloader, nil
}

func (this *myDownloaderPool) Return(downloader Downloader) error {
  return this.pool.Return(downloader)
}

func (this *myDownloaderPool) Total() uint32 {
  return this.pool.Total()
}

func (this *myDownloaderPool) Used() uint32 {
  return this.pool.Used()
}