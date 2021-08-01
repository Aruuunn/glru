package glru_test

import (
	"testing"

	"github.com/ArunMurugan78/glru"
	"github.com/stretchr/testify/assert"
)

func TestNewGlru(t *testing.T) {
	cache := glru.New(glru.Config{
		MaxItems: 4,
	})

	val := struct{ Name string }{"arun"}
	cache.Set("One", val)
	cache.Set("Two", 2)
	cache.Set("Three", "3")
	cache.Set("Four", 4)

  got, err := cache.Get("One")
  
  assert.NotNil(t, err)

	assert.Equal(t, val, got)
}
