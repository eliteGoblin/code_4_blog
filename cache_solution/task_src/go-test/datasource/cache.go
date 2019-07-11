package datasource

import (
	"time"
)

type DistributedCache struct {
	data map[string]interface{}
}

func (dc *DistributedCache) Value(key string) (interface{}, error) {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	return dc.data[key], nil
}

func (dc *DistributedCache) Store(key string, value interface{}) error {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	dc.data[key] = value
	return nil
}
