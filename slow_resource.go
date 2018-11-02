package cache

import (
	"errors"
	"fmt"
	"time"
)

/**
slowResource simulates a resources with high latency such as i/o, network,...
*/
type slowResource struct {
	store      map[string]string
	getLatency int
	putLatency int
}

func NewSlowResource(getLatency, putLatency int) *slowResource {
	return &slowResource{
		store:      make(map[string]string),
		getLatency: getLatency,
		putLatency: putLatency,
	}
}

func (r *slowResource) Get(key string) (interface{}, error) {
	<-time.After(time.Duration(r.getLatency) * time.Millisecond)
	if v, ok := r.store[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Key not exists %s", key)
}

func (r *slowResource) Put(key string, val interface{}) error {
	<-time.After(time.Duration(r.putLatency) * time.Millisecond)
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
