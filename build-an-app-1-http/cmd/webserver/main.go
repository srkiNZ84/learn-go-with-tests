package main

import (
	"github.com/srkiNZ84/learn-go-with-tests/build-an-app-1-http"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	//store := NewInMemoryPlayerStore()
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening db %s, %v", dbFileName, err)
	}
	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating filesystem player store, %v", err)
	}
	server := poker.NewPlayerServer(store)
	
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
