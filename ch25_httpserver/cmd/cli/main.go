package main

import (
	"bytes"
	poker "ch25/httpserver"
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)
	poker.NewCLI(os.Stdin, &bytes.Buffer{}, game).PlayPoker()
}
