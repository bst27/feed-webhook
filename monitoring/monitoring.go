package monitoring

import (
	"context"
	"fmt"
	"github.com/bst27/feed-webhook/reader"
	"github.com/bst27/feed-webhook/webhook"
	"log"
	"net/url"
	"time"
)

type Monitoring interface {
	AddFeed(url string, pollingInterval int64, reader reader.Reader, webhooks []webhook.Receiver)
	Start(ctx context.Context)
}

func Create() Monitoring {
	return &monitoring{}
}

type monitoring struct {
	feeds []*feed
}

func (m *monitoring) Start(ctx context.Context) {
	for _, f := range m.feeds {
		go m.monitor(ctx, f)
	}
}
func (m *monitoring) monitor(ctx context.Context, f *feed) {
	for {
		if f.getNextPollingTime().Before(time.Now()) {
			fmt.Println(time.Now(), "Checking", f.url.String())
			err := f.CheckChanges()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(time.Now(), "Checked", f.url.String())
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Until(f.getNextPollingTime())):
		}
	}
}

func (m *monitoring) AddFeed(feedUrl string, pollingInterval int64, reader reader.Reader, webhooks []webhook.Receiver) {
	uri, err := url.ParseRequestURI(feedUrl)
	if err != nil {
		log.Fatal(err)
	}

	interval, err := time.ParseDuration(fmt.Sprintf("%ds", pollingInterval))
	if err != nil {
		log.Fatal(err)
	}

	f := &feed{
		url:             *uri,
		pollingInterval: interval,
		lastPolling:     time.Time{},
		reader:          reader,
		webhooks:        webhooks,
	}

	m.feeds = append(m.feeds, f)
}
