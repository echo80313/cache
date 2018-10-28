package cache

import (
	"errors"
	"fmt"
	"time"
)

type slowResource struct {
	store map[string]string
}

func NewSlowResource() *slowResource {
	return &slowResource{
		store: make(map[string]string),
	}
}

func (r *slowResource) Get(key string) (interface{}, error) {
	<-time.After(50 * time.Millisecond)
	if v, ok := r.store[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Key not exists %s", key)
}

func (r *slowResource) Put(key string, val interface{}) error {
	<-time.After(250 * time.Millisecond)
	return r.FastPut(key, val)
}

func (r *slowResource) FastPut(key string, val interface{}) error {
	if v, ok := val.(string); ok {
		r.store[key] = v
		return nil
	}
	return errors.New("Invalid value type")
}

var _ Resource = &slowResource{}
