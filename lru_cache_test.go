package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCacheListAddAndAccess(t *testing.T) {
	l := newLruCacheNodeList(100)
	kvSet := genKeyValueSet(10, 5, 8)
	for k, v := range kvSet {
		l.Add(k, v)
		val, ok := l.Access(k)
		assert.True(t, ok)
		assert.Equal(t, v, val)
	}
	assert.Equal(t, 10, l.getSize())

	l = newLruCacheNodeList(1)
	for k, v := range kvSet {
		l.Add(k, v)
		val, ok := l.Access(k)
		assert.True(t, ok)
		assert.Equal(t, v, val)
	}

	assert.Equal(t, 1, l.getSize())

	key := "non-exist"
	_, ok := l.Access(key)
	assert.False(t, ok)
}

func TestLRUCacheListAccessMoveToHead(t *testing.T) {
	keySet := []string{"k1", "k2", "k3"}
	val := "random value" // choose by rolling a dice

	l := newLruCacheNodeList(2)
	l.Add(keySet[0], val)
	l.Add(keySet[1], val)
	l.Add(keySet[2], val)

	_, ok := l.Access(keySet[0])
	assert.False(t, ok) // k1 is purged
	_, ok = l.Access(keySet[1])
	assert.True(t, ok)
	_, ok = l.Access(keySet[2])
	assert.True(t, ok)

	l.Access(keySet[1]) // touch k2
	l.Add(keySet[0], val)
	_, ok = l.Access(keySet[0])
	assert.True(t, ok)
	_, ok = l.Access(keySet[1])
	assert.True(t, ok)
	_, ok = l.Access(keySet[2])
	assert.False(t, ok) // k3 is purged
}
