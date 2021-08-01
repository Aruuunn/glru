package glru


type Glru struct {}


type Config struct {
  MaxItems int
}


func New(config Config) *Glru {
  return &Glru{}
}

func (cache *Glru) Set(key string, value interface{}) {}


func (cache *Glru) Get(key string) interface{} {
  return nil
}
