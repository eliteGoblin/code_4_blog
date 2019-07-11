package inmemcache

import (
	"container/list"
	"sync"

	errors "golang.org/x/xerrors"
)

const (
	InMemCacheCapacity = 1000
)

type LRUCache struct {
	mapKV    map[string]*list.Element
	list     *list.List
	capacity int
	mutex    sync.Mutex
}

type element struct {
	key   string
	value interface{}
}

func New(capacity int) (cache *LRUCache, err error) {
	if capacity <= 0 {
		return &LRUCache{}, errors.New("invalid capacity number")
	}
	cache = &LRUCache{}
	cache.mapKV = make(map[string]*list.Element)
	cache.capacity = capacity
	cache.list = list.New()
	return cache, nil
}

func (cache *LRUCache) Value(key string) (value interface{}, err error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if node, ok := cache.mapKV[key]; ok {
		e, ok := node.Value.(*element)
		if !ok {
			return 0, errors.Errorf("invalid value type in Value: %+v", node.Value)
		}
		cache.list.MoveToFront(node)
		return e.value, nil
	}
	return nil, nil
}

func (cache *LRUCache) Store(key string, value interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	// move recently accessed element to front
	if node, ok := cache.mapKV[key]; ok {
		node.Value = cache.makeNode(key, value)
		cache.list.MoveToFront(node)
		return nil
	}
	var node *list.Element
	if cache.capacity > cache.list.Len() {
		node = cache.list.PushFront(cache.makeNode(key, value))
	} else {
		// remove least recent used element(tail of list)
		tail := cache.list.Back()
		if tail == nil {
			return errors.New("capacity of cache is 0")
		}
		e, ok := tail.Value.(*element)
		if !ok {
			return errors.Errorf("invalid value type in Store: %+v", tail.Value)
		}
		cache.list.Remove(tail)
		delete(cache.mapKV, e.key)
		node = cache.list.PushFront(cache.makeNode(key, value))
	}
	cache.mapKV[key] = node
	return nil
}

func (cache *LRUCache) makeNode(key string, value interface{}) *element {
	return &element{
		key:   key,
		value: value,
	}
}
