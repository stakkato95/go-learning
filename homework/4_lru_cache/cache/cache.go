package cache

import "github.com/stakkato95/lru_cache/mylist"

type LruCache interface {
	Set(key string, value mylist.Value) bool
	Get(key string) (mylist.Value, bool)
	Clear()
}

type lruCache struct {
	capacity int
	keys     map[string]*mylist.Item
	list     mylist.List
}

func NewCache(cap int) LruCache {
	return &lruCache{capacity: cap, keys: map[string]*mylist.Item{}, list: mylist.NewList()}
}

func (l *lruCache) Set(key string, value mylist.Value) bool {
	//item is already in list
	if item, ok := l.keys[key]; ok {
		l.list.MoveToFront(item)
		return true
	}

	//item is not in list
	item := l.list.PushFront(value)
	l.keys[key] = item

	//capacity is exceeded
	if l.capacity < l.list.Len() {
		evictedItem := l.list.Back()
		l.list.Remove(evictedItem)
		for evictedKey, val := range l.keys {
			if val == evictedItem {
				delete(l.keys, evictedKey)
				break
			}
		}
	}

	return false
}

func (l *lruCache) Get(key string) (mylist.Value, bool) {
	if item, ok := l.keys[key]; ok {
		l.keys[key] = l.list.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.keys = map[string]*mylist.Item{}
	l.list.Clear()
}
