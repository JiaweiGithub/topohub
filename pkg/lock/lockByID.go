package lock

import (
	"sync"
)

// LockManager manages locks for different IDs
type LockManager struct {
	locks sync.Map
}

// Exported instance of LockManager
var LockManagerInstance = &LockManager{}

// GetLock retrieves a lock for the given ID, creating one if it doesn't exist
func (lm *LockManager) GetLock(name string) *sync.Mutex {
	if lock, ok := lm.locks.Load(name); ok {
		return lock.(*sync.Mutex)
	}

	// Create a new lock
	newLock := &sync.Mutex{}
	lm.locks.Store(name, newLock)
	return newLock
}

// MyObject represents an object with an ID
type MyObject struct {
	ID int
}
