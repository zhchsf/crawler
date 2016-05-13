package pool

import(
  "errors"
  "reflect"
  "sync"
)

type Entity interface {
  Id() uint32
}

type Pool interface {
  Take() (Entity, error)
  Return(entity Entity) error
  Total() uint32
  Used() uint32
}

type myPool struct {
  total uint32
  eType reflect.Type
  genEntity func() Entity
  container chan Entity
  idContainer map[uint32]bool
  mutex sync.Mutex
}

func NewPool(total uint32, entityType reflect.Type, genEntity func() Entity) (Pool, error){
  if total == 0 {
    return nil, errors.New("初始化pool size需大于0")
  }

  size := int(total)
  container := make(chan Entity, size)
  idContainer := make(map[uint32]bool)
  for i := 0; i < size; i++ {
    newEntity := genEntity()
    if entityType != reflect.TypeOf(newEntity) {
      return nil, errors.New("entity类型错误")
    }
    container <- newEntity
    idContainer[newEntity.Id()] = true
  }

  pool := &myPool{
    total: total, 
    eType: entityType, 
    genEntity: genEntity, 
    container: container, 
    idContainer: idContainer,
  }
  return pool, nil
}

func (this *myPool) Take() (Entity, error) {
  entity := <-this.container
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.idContainer[entity.Id()] = false
  return entity, nil
}

func (this *myPool) Return(entity Entity) error {
  if entity == nil || reflect.TypeOf(entity) != this.eType {
    return errors.New("entity对象为nil或者类型错误")
  }
  if _, ok := this.idContainer[entity.Id()]; !ok {
    return errors.New("归还entity不存在")
  }
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.idContainer[entity.Id()] = true
  this.container <- entity
  return nil
}

func (this *myPool) Total() uint32 {
  return this.total
}

func (this *myPool) Used() uint32 {
  return this.total - uint32(len(this.container))
}
