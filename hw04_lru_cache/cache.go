package hw04lrucache

import "sync"

type Key string

type QueueItem struct {
	Key   Key
	Value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	item, ok := cache.items[key]

	queueItem := QueueItem{
		key,
		value,
	}

	if !ok {
		if cache.queue.Len() == cache.capacity {
			lastEl := cache.queue.Back()
			lastElKey := lastEl.Value.(QueueItem).Key
			cache.queue.Remove(lastEl)
			delete(cache.items, lastElKey)
		}

		cache.items[key] = cache.queue.PushFront(queueItem)
	} else {
		item.Value = queueItem
		cache.queue.MoveToFront(item)
	}

	return ok
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	item, ok := cache.items[key]

	if !ok {
		return nil, false
	}
	cache.queue.MoveToFront(item)

	return item.Value.(QueueItem).Value, true
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
	cache.mutex.Unlock()
}
