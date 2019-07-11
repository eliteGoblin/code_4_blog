package datasource

import "time"

type Database struct {
	data map[string]interface{}
}

func (db *Database) Value(key string) (interface{}, error) {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)

	return db.data[key], nil
}

func (db *Database) Store(key string, value interface{}) error {
	// simulate 500ms roundtrip to the distributed cache
	time.Sleep(500 * time.Millisecond)

	db.data[key] = value
	return nil
}
