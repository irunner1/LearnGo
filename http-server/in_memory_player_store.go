package main

// NewInMemoryPlayerStore creates new player
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore stores player score in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin calls POST method to record player score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore gets player score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
