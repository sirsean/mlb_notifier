package main

import (
	"github.com/sirsean/mlb_notifier/poll"
	"log"
)

func main() {
	log.Println("This is MLB Notifier, starting up")
	poll.Start()
}
