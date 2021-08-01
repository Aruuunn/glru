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

	assertNotFound := func(key string) {
		_, err := cache.Get(key)

		assert.Equal(t, glru.ErrKeyNotFound, err)
	}

	assertExpectedValue := func(key string, expectedValue interface{}) {
		val, err := cache.Get(key)
		assert.Nil(t, err)
		assert.Equal(t, expectedValue, val)
	}

	val := struct{ Name string }{"arun"}
	cache.Set("One", val)
	cache.Set("Two", 2)
	cache.Set("Three", "3")
	cache.Set("Four", 4)

	assertExpectedValue("One", val)

	assertNotFound("KeyNotPresent")

	cache.Set("Five", 5)

	assertNotFound("Two")

	cache.Set("Six", "6")

	assertNotFound("Three")

	assertExpectedValue("Six", "6")

	assertExpectedValue("Five", 5)

	assertExpectedValue("Four", 4)

	cache.Set("Name", "arun")
	cache.Set("Golang", "Awesome")

	assertNotFound("Six")
	assertExpectedValue("Name", "arun")
}
