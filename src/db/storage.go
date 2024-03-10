package db

import "sync"

type Storage interface {
	CreateAccount() error
}

var storage Storage

var lock = &sync.Mutex{}

// implements singleton pattern to always get same storage
func GetStorage() Storage {
	if storage == nil {
		lock.Lock()
		defer lock.Unlock()
		if storage == nil {
			// creates storage instance
			storage = newPostgressStoreOrFatal()
		}
	}
	return storage
}