package main

import (
	"log"
	"os"
	"github.com/ryanwolhuter/test-driven-web-app"
	"fmt"
)

const dbFileName = "game.db.json"

func main() {
    store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Let's play poker")
    fmt.Println("Type {Name} wins to record a win")
    poker.NewCLI(store, os.Stdin).PlayPoker()
}