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
	mutex    sync.RWMutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.RLock()
	item, ok := cache.items[key]
	cache.mutex.RUnlock()

	if !ok {
		cache.mutex.Lock()
		if cache.queue.Len() == cache.capacity {
			lastEl := cache.queue.Back()
			lastElKey := lastEl.Value.(QueueItem).Key
			cache.queue.Remove(lastEl)
			delete(cache.items, lastElKey)
		}
		cache.mutex.Unlock()

		queueItem := QueueItem{
			key,
			value,
		}

		cache.mutex.Lock()
		cache.items[key] = cache.queue.PushFront(queueItem)
		cache.mutex.Unlock()
	} else {
		item.Value = QueueItem{
			key,
			value,
		}

		cache.mutex.Lock()
		cache.queue.MoveToFront(item)
		cache.mutex.Unlock()
	}

	return ok
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.RLock()
	item, ok := cache.items[key]
	cache.mutex.RUnlock()

	if !ok {
		return nil, false
	}

	cache.mutex.Lock()
	cache.queue.MoveToFront(item)
	cache.mutex.Unlock()

	return item.Value.(QueueItem).Value, true
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
	cache.mutex.Unlock()
}
