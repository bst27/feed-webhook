package config

import (
	"github.com/google/uuid"
	"net/url"
	"time"
)

type Config struct {
	Feeds    []*Feed
	Webhooks []*Webhook
}

func (c *Config) AddFeed(url url.URL, feedType string, interval time.Duration) {
	for _, f := range c.Feeds {
		if f.URL == url.String() {
			f.PollingInterval = int64(interval.Seconds())
			return
		}
	}

	c.Feeds = append(c.Feeds, &Feed{
		ID:              uuid.New().String(),
		URL:             url.String(),
		PollingInterval: int64(interval.Seconds()),
		Type:            feedType,
	})
}

func (c *Config) HasFeed(feedId string) bool {
	for _, f := range c.Feeds {
		if f.ID == feedId {
			return true
		}
	}

	return false
}

func (c *Config) GetWebhooks(feedId string) []*Webhook {
	webhooks := make([]*Webhook, 0)

	for _, w := range c.Webhooks {
		if w.FeedID == feedId {
			webhooks = append(webhooks, w)
		}
	}

	return webhooks
}

func (c *Config) AddWebHook(feedId string, url string) {
	c.Webhooks = append(c.Webhooks, &Webhook{
		ID:     uuid.New().String(),
		FeedID: feedId,
		URL:    url,
		Type:   "web",
	})
}

func (c *Config) AddFilesystemHook(feedId string, directory string) {
	c.Webhooks = append(c.Webhooks, &Webhook{
		ID:     uuid.New().String(),
		FeedID: feedId,
		URL:    directory,
		Type:   "filesystem",
	})
}

func (c *Config) AddStdoutHook(feedId string) {
	c.Webhooks = append(c.Webhooks, &Webhook{
		ID:     uuid.New().String(),
		FeedID: feedId,
		URL:    "",
		Type:   "stdout",
	})
}
