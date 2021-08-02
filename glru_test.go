package glru_test

import (
	"log"
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
	t.Run("Test Case 1", func(t *testing.T) {
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
		assertExpectedValue("Golang", "Awesome")
		assertExpectedValue("Five", 5)

		assertExpectedValue("Four", 4)

		cache.Set("TCE", "Sucks")
		assertNotFound("Name")
	})

	testLRU(t, 2, map[string]interface{}{
		"1":    6,
		"6":    7,
		"&s":   "dsnfnkd",
		"name": "arun",
	}, []string{"6", "1", "&s", "name", "&s", "1", "&s", "6", "1", "&s", "name", "name", "1", "&s"},
		[]bool{true, true, true, true, false, true, false, true, true, true, true, false, true, true})

	testLRU(t, 3, map[string]interface{}{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
		"Eight": 8,
	}, []string{"One", "Two", "Three", "Four", "One", "Three", "Two", "Four", "One", "Two", "Four"},
		[]bool{true, true, true, true, true, false, true, true, true, false, false})
}

func testLRU(t *testing.T, maxItems int, keyValueMap map[string]interface{}, refSequence []string, faults []bool) {
	cache := glru.New(glru.Config{MaxItems: maxItems})

	for idx, ref := range refSequence {
		val, err := cache.Get(ref)

		assert.Equal(t, faults[idx], err == glru.ErrKeyNotFound)

		if err == nil {
			assert.Equal(t, keyValueMap[ref], val)
		} else if err == glru.ErrKeyNotFound {
			cache.Set(ref, keyValueMap[ref])
		} else {
			log.Fatalln("Unexpected error ", err.Error())
		}
	}
}
