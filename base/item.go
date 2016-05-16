package base

type Item map[string]interface{}

func (this *Item) Valid() bool {
  return this != nil
}
