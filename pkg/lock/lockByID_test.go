package lock

import "testing"

// TestGetLock tests the GetLock method of LockManager
func TestGetLock(t *testing.T) {
	lm := &LockManager{}

	// Test getting a lock for a specific ID
	lock1 := lm.GetLock("host1")
	lock2 := lm.GetLock("host1")

	// Ensure both locks are the same
	if lock1 != lock2 {
		t.Errorf("Expected the same lock for ID 1")
	}

	// Test getting a lock for a different ID
	lock3 := lm.GetLock("host2")

	// Ensure lock3 is different from lock1
	if lock1 == lock3 {
		t.Errorf("Expected different locks for ID 1 and ID 2")
	}
}
