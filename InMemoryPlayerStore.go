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

// GetLeague returns a slice of Player, which represents a league.
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}