package cache

type LRUCache struct {
	resource Resource
	cache    []interface{}
}

func NewLRUCache(resource Resource, sz int) *LRUCache {
	return &LRUCache{
		resource: resource,
		cache:    make([]interface{}, sz),
	}
}

func (c *LRUCache) Get(key string) interface{} {
	return nil
}

func (c *LRUCache) Put(key string, val interface{}) error {
	return nil
}

// make sure cache is kind of resource
var _ Resource = &LRUCache{}
