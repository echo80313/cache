package cache

type Resource interface {
	Get(string) interface{}
	Put(string, interface{}) error
}

const (
	CacheTypeNon = 0
	CacheTypeLRU = 1
)
