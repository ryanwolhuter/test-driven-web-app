package main

import (
	"fmt"
	"encoding/json"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("Problem parsing league, %v", err)
	}

	return league, err
}