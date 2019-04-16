package main

//NewInMemoryPlayerStore initializes an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin records when a player wins.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore retrieves the specified player's score.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}