package main

import (
	"log"
	"os"

	db "github.com/jacobd39/admsys/bd"
	hnd "github.com/jacobd39/admsys/handlers"
)

func main() {
	os.Setenv("PORT", "8090")

	if db.CheckDBConnection() == 0 {
		log.Fatal("No DB connection")
		return
	}

	hnd.BackendHandlers()
}
