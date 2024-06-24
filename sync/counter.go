package sync

import "sync"

type Counter struct {
	mu	sync.Mutex	// Locks
	value	int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

/* 
	LESSONS LEARNED: 
	- Use channels when you're passing ownership of data (use `go race` to check race conditions)
	- Use mutexes to manage state (use `go vet` to check for subtle bugs)
*/