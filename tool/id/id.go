package id_generator

import(
  "math"
  "sync"
)

type IdGenerator interface {
  NextId() uint32
}

type myIdGenerator struct {
  id uint32
  finished bool
  mutex sync.Mutex
}

func NewIdGenerator() IdGenerator {
  return &myIdGenerator{}
}

func (this *myIdGenerator) NextId() uint32 {
  id := this.id
  this.mutex.Lock()
  defer this.mutex.Unlock()
  if this.finished {
    defer func() {
      this.finished = false
      this.id = 0
    }()
  }
  this.id++
  if this.id >= math.MaxUint32 {
    this.finished = true
  }
  return id
}
