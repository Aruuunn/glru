// package glru encapsulates logic for the LRU cache.
package glru

import (
	"errors"
	"sync"

	"github.com/ArunMurugan78/glru/dll"
)

const DEFAULT_MAX_ITEMS = 10000

// Glru is the main struct which implements the LRU cache.
type Glru struct {
	nodeMap  map[string]*dll.Node
	maxItems int
	list     *dll.Dll
	items    int
}

// Config is passed to New().
type Config struct {
	MaxItems int // Defaults to DEFAULT_MAX_ITEMS
}

// ErrKeyNotFound is returned by Get method when the key is not found.
var ErrKeyNotFound = errors.New("key not found")

// New returns a new initialized Glru instance.
func New(config Config) *Glru {
	if config.MaxItems <= 0 {
		config.MaxItems = DEFAULT_MAX_ITEMS
	}

	return &Glru{maxItems: config.MaxItems, list: dll.New(), nodeMap: make(map[string]*dll.Node)}
}

// Set adds the key-value pair to the cache.
func (cache *Glru) Set(key string, value interface{}) {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	node, ok := cache.nodeMap[key]

	if ok {
		if node.Value != value {
			node.Value = value
		}

		return
	}

	if cache.items == cache.maxItems {
		// Cache is full, So delete Least Recently Used
		lastNode := cache.list.GetTail()

		cache.list.DeleteNode(lastNode)
		delete(cache.nodeMap, lastNode.Key)

		lastNode.Prev = nil
		lastNode.Next = nil

		cache.items--
	}

	cache.items++
	cache.nodeMap[key] = cache.list.Prepend(key, value)
}

// Get returns the value association with the key. Returns ErrKeyNotFound if the key is not found in cache.
func (cache *Glru) Get(key string) (interface{}, error) {
	var mutex sync.Mutex

	node, ok := cache.nodeMap[key]

	if ok {
		mutex.Lock()
		defer mutex.Unlock()

		// Brings the accessed node to the front of the list in O(1) time complexity
		cache.list.DeleteAndInsertAtHead(node)
		return node.Value, nil
	}
	return nil, ErrKeyNotFound
}

