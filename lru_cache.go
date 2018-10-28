package cache

import (
	"errors"
)

type cacheNode struct {
	key  string
	data interface{}
	next *cacheNode
	prev *cacheNode
}

type lruCacheNodeList struct {
	head *cacheNode
	tail *cacheNode

	// Store the key -> cachenode mapping.
	index map[string]*cacheNode

	cap int
}

func newLruCacheNodeList(cap int) *lruCacheNodeList {
	head := &cacheNode{} // sentinal node
	tail := &cacheNode{} // tail sentinal node
	head.prev = head
	head.next = tail
	tail.prev = head
	tail.next = tail
	return &lruCacheNodeList{
		cap:   cap,
		head:  head,
		tail:  tail,
		index: make(map[string]*cacheNode),
	}
}

func (l *lruCacheNodeList) Access(key string) (interface{}, bool) {
	if node, ok := l.index[key]; ok {
		l.moveToHead(node)
		return node.data, true
	}
	return nil, false
}

// Add assumes that key is not in cache, or it would override what we
// have in cache
func (l *lruCacheNodeList) Add(key string, val interface{}) {
	node := &cacheNode{
		key:  key,
		data: val,
	}
	l.index[key] = node
	l.AddToHead(node)
	if l.getSize() > l.cap {
		l.popTail()
	}
}

func (l *lruCacheNodeList) removeNode(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *lruCacheNodeList) moveToHead(node *cacheNode) {
	l.removeNode(node)
	l.AddToHead(node)
}

func (l *lruCacheNodeList) popTail() {
	nodeToPurge := l.tail.prev
	l.removeNode(nodeToPurge)
	delete(l.index, nodeToPurge.key)
	l.tail.prev = nodeToPurge.prev // deference so that GC can clean it up
}

func (l *lruCacheNodeList) getSize() int {
	return len(l.index)
}

func (l *lruCacheNodeList) AddToHead(node *cacheNode) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

type LRUCache struct {
	resource Resource

	// memory limit
	capacity int
	cache    *lruCacheNodeList
}

type metaData struct {
	where int
}

func NewLRUCache(resource Resource, sz int) *LRUCache {
	return &LRUCache{
		resource: resource,
	}
}

func (c *LRUCache) Get(key string) (interface{}, error) {
	if val, ok := c.cache.Access(key); ok {
		return val, nil
	}

	// cache miss, sad :(
	val, err := c.resource.Get(key)
	if err == nil {
		c.cache.Add(key, val)
	}
	return val, err
}

func (c *LRUCache) Put(key string, val interface{}) error {
	return errors.New("Not implemented")
}

// make sure cache is kind of resource
var _ Resource = &LRUCache{}
