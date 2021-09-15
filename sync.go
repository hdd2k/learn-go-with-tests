package main

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// lock with mutex (mutual exclusion lock) to have ONLY 1 process access at a time
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}
