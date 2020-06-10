package monitoring

import (
	"github.com/bst27/feed-webhook/reader"
	"github.com/bst27/feed-webhook/webhook"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type feed struct {
	url             url.URL
	pollingInterval time.Duration
	lastPolling     time.Time
	reader          reader.Reader
	webhooks        []webhook.Receiver
	lastContent     []byte
}

func (f *feed) isPollingDue() bool {
	return time.Now().Sub(f.lastPolling) >= f.pollingInterval
}

func (f *feed) getNextPollingTime() time.Time {
	return f.lastPolling.Add(f.pollingInterval)
}

func (f *feed) hasChanged() (bool, error) {
	content, err := f.getContent()
	if err != nil {
		return false, err
	}

	return f.reader.HasChanged(content, f.lastContent)
}

func (f *feed) getContent() ([]byte, error) {
	resp, err := http.Get(f.url.String())
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	f.lastPolling = time.Now()

	return content, nil
}

func (f *feed) CheckChanges() error {
	changed, err := f.hasChanged()
	if err != nil {
		return err
	}

	if changed {
		content, err := f.getContent()
		if err != nil {
			return err
		}

		changes, err := f.reader.GetChanges(content, f.lastContent)
		if err != nil {
			return err
		}

		// avoid sending webhooks on first iteration (init)
		if f.lastContent != nil {
			for _, change := range changes {
				for _, w := range f.webhooks {
					err = w.Receive(change)
				}
			}
		}
		f.lastContent = content

		return nil
	}

	return nil
}
