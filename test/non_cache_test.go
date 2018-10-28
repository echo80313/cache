package test

import (
	"cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonCachePutAndGet(t *testing.T) {
	c := cache.NewNonCache(NewSlowResource())
	key := "key1"
	value := "val1"
	c.Put(key, value)
	assert.Equal(t, c.Get(key), value)

	// Get nonexist-key
	assert.Equal(t, c.Get("Non-exist"), "")
}

func BenchmarkGet1(b *testing.B) {
	benchmarkGet(1, b)
}
func BenchmarkGet10(b *testing.B) {
	benchmarkGet(10, b)
}

func BenchmarkGet100(b *testing.B) {
	benchmarkGet(100, b)
}

func benchmarkGet(n int, b *testing.B) {
	res := NewSlowResource()
	key := "key1"
	value := "val1"
	res.FastPut(key, value)

	c := cache.NewNonCache(res)
	for i := 0; i < n; i++ {
		c.Get(key)
	}
}

func BenchmarkPutAndGet1(b *testing.B) {
	benchmarkPutAndGet(GenRandomOps(1), b)
}

func BenchmarkPutAndGet10(b *testing.B) {
	benchmarkPutAndGet(GenRandomOps(10), b)
}

func BenchmarkPutAndGet100(b *testing.B) {
	benchmarkPutAndGet(GenRandomOps(100), b)
}

func benchmarkPutAndGet(ops []*Op, b *testing.B) {
	c := cache.NewNonCache(NewSlowResource())
	for _, op := range ops {
		switch op.t {
		case GetOp:
			c.Get(op.key)
		case PutOp:
			c.Put(op.key, op.val)
		}
	}
}
