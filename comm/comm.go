package comm

import (
	"errors"
	twilio "github.com/carlosdp/twiliogo"
	"github.com/sirsean/go-mailgun/mailgun"
	"github.com/sirsean/mlb_notifier/event"
	"log"
	"os"
	"sync"
	"time"
)

var emailTo string = os.Getenv("EMAIL_TO")
var emailFrom string = os.Getenv("EMAIL_FROM")
var mailgunApiDomain string = os.Getenv("MAILGUN_API_DOMAIN")
var mailgunApiKey string = os.Getenv("MAILGUN_API_KEY")

var mailClient *mailgun.Client

var twilioAccountSid string = os.Getenv("TWILIO_ACCOUNT_SID")
var twilioAuthToken string = os.Getenv("TWILIO_AUTH_TOKEN")
var smsTo string = os.Getenv("SMS_TO")
var smsFrom string = os.Getenv("SMS_FROM")

var twilioClient twilio.Client

func init() {
	if mailgunApiKey == "" && twilioAccountSid == "" {
		log.Fatal("Either MAILGUN_API_KEY or TWILIO_ACCOUNT_SID is required")
	}

	if mailgunApiKey != "" {
		if emailTo == "" {
			log.Fatal("EMAIL_TO is required")
		}
		if emailFrom == "" {
			log.Fatal("EMAIL_FROM is required")
		}
		if mailgunApiDomain == "" {
			log.Fatal("MAILGUN_API_DOMAIN is required")
		}
		if mailgunApiKey == "" {
			log.Fatal("MAILGUN_API_KEY is required")
		}

		mailClient = mailgun.NewClient(mailgunApiKey, mailgunApiDomain)
	}

	if twilioAccountSid != "" {
		if twilioAuthToken == "" {
			log.Fatal("TWILIO_AUTH_TOKEN is required")
		}
		if smsTo == "" {
			log.Fatal("SMS_TO is required")
		}
		if smsFrom == "" {
			log.Fatal("SMS_FROM is required")
		}

		twilioClient = twilio.NewClient(twilioAccountSid, twilioAuthToken)
	}
}

// store a map of date keys ("yyyy/mm/dd") to a list of events
// that were sent for games on that day. This way we can make
// sure that we don't send duplicate events.
var store = struct {
	m map[string][]event.Event
	sync.Mutex
}{m: make(map[string][]event.Event)}

func Send(events ...event.Event) {
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
			go sendSms(e)
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
	log.Printf("Clean comm records, keep: %v, %v\n", todayKey, yesterdayKey)
	for k, _ := range store.m {
		if k != todayKey && k != yesterdayKey {
			log.Printf("Delete comm records: %v\n", k)
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
	if mailClient != nil {
		log.Printf("Sending Event: %v\n", e)
		mailClient.Send(mailgun.Message{
			FromName:    "MLB Notifier",
			FromAddress: emailFrom,
			ToAddress:   emailTo,
			Subject:     e.Summary(),
			Body:        e.Summary(),
		})
	}
}

func sendSms(e event.Event) error {
	if twilioClient != nil {
		_, err := twilio.NewMessage(twilioClient, smsFrom, smsTo, twilio.Body(e.Summary()))
		return err
	} else {
		return errors.New("No twilioClient")
	}
}
