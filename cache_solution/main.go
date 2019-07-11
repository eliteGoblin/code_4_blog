package main

import (
	"cache_solution/datasource"
	"cache_solution/infra/cache"
	"cache_solution/infra/database"
	"cache_solution/infra/inmemcache"
	"cache_solution/pkg/keygen"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	db := database.New()
	database.FillDatabase(db)
	logrus.Info("fill database completed")
	local, err := inmemcache.New(inmemcache.InMemCacheCapacity)
	if err != nil {
		logrus.Errorf("%+v", err)
		return
	}
	external := cache.New()
	// create retriever to get key from cache
	retriever := datasource.NewDataRetrieve(local, external, db)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			var startTime time.Time
			for j := 0; j < 50; j++ {
				startTime = time.Now()
				key := keygen.RandomKey(0, 9)
				value, err := retriever.Value(key)
				if err != nil {
					logrus.Errorf("%+v", err)
					continue
				}
				logrus.Infof("Request '%s', response '%s', time: %.2f ms",
					key, value.(string), float64(time.Now().Sub(startTime))/float64(time.Millisecond))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
