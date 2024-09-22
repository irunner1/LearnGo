package main

import "sync"

// Counter count numbers
type Counter struct {
	value int
	mu    sync.Mutex
}

// Inc increments number
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value gets value of couner
func (c *Counter) Value() int {
	return c.value
}
