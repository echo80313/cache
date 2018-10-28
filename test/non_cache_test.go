package test

import (
	"cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonCachePutAndGet(t *testing.T) {
	c := cache.NewNonCache(NewSlowResource(), 100)
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

func BenchmarkGet1000(b *testing.B) {
	benchmarkGet(1000, b)
}

func benchmarkGet(n int, b *testing.B) {
	res := NewSlowResource()
	key := "key1"
	value := "val1"
	res.FastPut(key, value)

	c := cache.NewNonCache(res, 100)
	for i := 0; i < n; i++ {
		c.Get(key)
	}
}
