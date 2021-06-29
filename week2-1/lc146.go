package week21

import "container/list"

// https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	innerMap  map[int]*list.Element
	innerData *list.List
	cap       int
}

type LRUCacheEntry struct {
	key, value int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		innerMap:  make(map[int]*list.Element, capacity),
		cap:       capacity,
		innerData: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if vptr, ok := c.innerMap[key]; ok {
		c.innerData.MoveToFront(vptr)
		return vptr.Value.(*LRUCacheEntry).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if vptr, ok := c.innerMap[key]; ok {
		c.innerData.MoveToFront(vptr)
		vptr.Value.(*LRUCacheEntry).value = value
	} else {
		node := &LRUCacheEntry{key: key, value: value}
		c.innerMap[key] = c.innerData.PushFront(node)
		if len(c.innerMap) > c.cap {
			last := c.innerData.Back()
			c.innerData.Remove(last)
			delete(c.innerMap, last.Value.(*LRUCacheEntry).key)
		}
	}
}
