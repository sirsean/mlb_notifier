package main

import (
	"fmt"
	"github.com/sirsean/go-mailgun/mailgun"
	"github.com/sirsean/mlb_notifier/comm"
	"github.com/sirsean/mlb_notifier/config"
	"github.com/sirsean/mlb_notifier/poll"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("This is MLB Notifier, starting up")
	config.LoadFile(filepath.Join(os.Getenv("HOME"), ".mlb_notifier"))
	comm.MailClient = mailgun.NewClient(config.Get("mailgun:api_key"), config.Get("mailgun:api_domain"))
	poll.Start()
}
