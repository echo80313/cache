package test

import (
	"errors"
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

func (r *slowResource) Get(key string) interface{} {
	<-time.After(50 * time.Millisecond)
	return r.store[key]
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
