package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/sirsean/mlb_notifier/poll"
    "github.com/sirsean/mlb_notifier/config"
    "github.com/sirsean/mlb_notifier/comm"
    "github.com/sirsean/go-mailgun/mailgun"
)

func main() {
    fmt.Println("This is MLB Notifier, starting up")
    config.LoadFile(filepath.Join(os.Getenv("HOME"), ".mlb_notifier"))
    //mailgun.ApiEndpoint = config.Get("mailgun:api_endpoint")
    //mailgun.ApiKey = config.Get("mailgun:api_key")
    comm.MailClient = mailgun.NewClient(config.Get("mailgun:api_key"), config.Get("mailgun:api_domain"))
    poll.Start()
}
