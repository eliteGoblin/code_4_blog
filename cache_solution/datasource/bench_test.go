package datasource

import (
	"cache_solution/infra/cache"
	"cache_solution/infra/database"
	"cache_solution/infra/inmemcache"
	"cache_solution/pkg/keygen"
	"testing"
)

func BenchmarkDataRetrieveValue(b *testing.B) {
	db := database.New()
	database.FillDatabase(db)
	external := cache.New()
	local, _ := inmemcache.New(inmemcache.InMemCacheCapacity)
	retriever := NewDataRetrieve(local, external, db)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			retriever.Value(keygen.RandomKey(0, 9))
		}
	})
}

func BenchmarkDataRetrieveValueLRU(b *testing.B) {
	db := database.New()
	database.FillDatabase(db)
	external := cache.New()
	local, _ := inmemcache.New(3)
	retriever := NewDataRetrieve(local, external, db)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			retriever.Value(keygen.RandomKey(0, 9))
		}
	})
}
