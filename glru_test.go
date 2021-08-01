package glru_test

import (
  "github.com/stretchr/testify/assert"
  "github.com/ArunMurugan78/glru"
  "testing"
)



func TestTestGlru(t *testing.T) {
  cache := glru.New(glru.Config {
    MaxItems: 4,
  })

  val := struct {Name string} {"arun"}
  cache.Set("One", val)
  cache.Set("Two", 2)
  cache.Set("Three", "3")
  cache.Set("Four", 4)

  assert.Equal(t, val, cache.Get("One"))
  assert.Equal(t, 2, cache.Get("One"))
  assert.Equal(t, "3", cache.Get("One"))
  assert.Equal(t, 4, cache.Get("One"))
}




