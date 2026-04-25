package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"naberi-backend/internal/naberi"
)

func main() {
	addr := os.Getenv("NABERI_ADDR")
	if addr == "" {
		addr = ":4010"
	}

	server := naberi.NewServer(naberi.NewStore(), time.Local)

	log.Printf("Naberi backend listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatal(err)
	}
}
