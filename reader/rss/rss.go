package rss

import (
	"encoding/xml"
)

func Get() *RSS {
	return &RSS{}
}

type RSS struct {
}

func (r RSS) HasChanged(feedContent []byte, prevFeedContent []byte) (bool, error) {
	return len(r.GetNewEntries(feedContent, prevFeedContent)) > 0, nil
}

func (r RSS) GetEntries(feedContent []byte) []item {
	f := &feed{}

	// Parsing errors are handled like no entries exist
	_ = xml.Unmarshal(feedContent, f)

	return f.Channel.Items
}

func (r RSS) GetNewEntries(feedContent []byte, prevFeedContent []byte) []item {
	currEntries := r.GetEntries(feedContent)
	prevEntries := make(map[string]item)
	for _, item := range r.GetEntries(prevFeedContent) {
		prevEntries[item.GUID] = item
	}

	changedEntries := make([]item, 0)
	for _, item := range currEntries {
		_, known := prevEntries[item.GUID]

		if !known {
			changedEntries = append(changedEntries, item)
		}
	}

	return changedEntries
}

func (r RSS) GetChanges(feedContent []byte, prevFeedContent []byte) ([]interface{}, error) {
	newEntries := r.GetNewEntries(feedContent, prevFeedContent)
	changes := make([]interface{}, 0)

	for _, item := range newEntries {
		changes = append(changes, item)
	}

	return changes, nil
}
