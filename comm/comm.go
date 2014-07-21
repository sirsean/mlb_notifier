package comm

import (
	"fmt"
	"github.com/sirsean/go-mailgun/mailgun"
	"github.com/sirsean/mlb_notifier/config"
	"github.com/sirsean/mlb_notifier/event"
	"sync"
	"time"
)

// store a map of date keys ("yyyy/mm/dd") to a list of events
// that were sent for games on that day. This way we can make
// sure that we don't send duplicate events.
var store = struct {
	m map[string][]event.Event
	sync.Mutex
}{m: make(map[string][]event.Event)}

var MailClient *mailgun.Client

func Send(events []event.Event) {
	for _, e := range events {
		dateKey := e.DateKey()
		store.Lock()
		_, ok := store.m[dateKey]
		if !ok {
			store.m[dateKey] = make([]event.Event, 0)
		}
		store.Unlock()
		if !isEventInStore(dateKey, e) {
			go sendEmail(e)
			store.Lock()
			store.m[dateKey] = append(store.m[dateKey], e)
			store.Unlock()
		}
	}
}

// Clean looks through the sent events and removes anything
// that was sent before yesterday. So, essentially, it keeps
// any date key that corresponds to today or yesterday.
func Clean() {
	todayKey, yesterdayKey := dateKeyFor(time.Now()), dateKeyFor(time.Now().Add(-24*time.Hour))
	fmt.Printf("Clean comm records, keep: %v, %v\n", todayKey, yesterdayKey)
	for k, _ := range store.m {
		if k != todayKey && k != yesterdayKey {
			fmt.Printf("Delete comm records: %v\n", k)
			store.Lock()
			delete(store.m, k)
			store.Unlock()
		}
	}
}

func dateKeyFor(t time.Time) string {
	year, month, day := t.Date()
	return event.DateKey(year, int(month), day)
}

func isEventInStore(dateKey string, event event.Event) bool {
	store.Lock()
	events, ok := store.m[dateKey]
	store.Unlock()
	if ok {
		for _, other := range events {
			if other == event {
				return true
			}
		}
	}
	return false
}

func sendEmail(e event.Event) {
	fmt.Printf("Sending Event: %v\n", e)
	MailClient.Send(mailgun.Message{
		FromName:    "MLB Notifier",
		FromAddress: config.Get("email:from"),
		ToAddress:   config.Get("email:to"),
		Subject:     e.Summary(),
		Body:        e.Summary(),
	})
}
