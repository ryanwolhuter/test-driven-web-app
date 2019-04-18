package main

import (
	"log"
	"os"
	"github.com/ryanwolhuter/test-driven-web-app"
	"fmt"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's player poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}