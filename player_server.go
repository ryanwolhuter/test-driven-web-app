package main

import (
	"fmt"
	"net/http"
)

// PlayerStore stores and fetches a player's score
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

//PlayerServer needs to be able to use a PlayerStore, so it needs a 
// reference to one. So we give PlayerServer an instance of PlayerStore.
//
// PlayerServer takes http.ResponseWriter, which implements
// io.Writer. This means we can use fmt.Fprint to send strings
// as HTTP responses.
type PlayerServer struct {
	store PlayerStore
}

//ServeHTTP is a method on the PlayerServer struct
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

// GetPlayerScore takes a player's name and returns their score.
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}