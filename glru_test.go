package glru_test

import (
	"reflect"
	"testing"

	"github.com/ArunMurugan78/glru"
	"github.com/stretchr/testify/assert"
)

func TestNewGlru(t *testing.T) {
	cache := glru.New(glru.Config{
		MaxItems: 4,
	})

	assert.NotNil(t, cache)
	assert.Equal(t, "*glru.Glru", reflect.TypeOf(cache).String())
}

func TestGetAndSet(t *testing.T) {
	cache := glru.New(glru.Config{
		MaxItems: 4,
	})

	val := struct{ Name string }{"arun"}
	cache.Set("One", val)
	cache.Set("Two", 2)
	cache.Set("Three", "3")
	cache.Set("Four", 4)

	got, err := cache.Get("One")

	assert.Nil(t, err)
	assert.Equal(t, val, got)

	_, err = cache.Get("KeyNotPresent")

	assert.Equal(t, err, glru.ErrKeyNotFound)

	cache.Set("Five", 5)

	_, err = cache.Get("One")

	assert.Equal(t, err, glru.ErrKeyNotFound)
}
