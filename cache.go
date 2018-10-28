package cache

type Resource interface {
	Get(string) (interface{}, error)
	Put(string, interface{}) error
}

const (
	CacheTypeNon = 0
	CacheTypeLRU = 1
)
