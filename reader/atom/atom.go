package atom

import (
	"encoding/xml"
)

func Get() *Atom {
	return &Atom{}
}

type Atom struct {
}

func (a Atom) HasChanged(feedContent []byte, prevFeedContent []byte) (bool, error) {
	return len(a.GetNewEntries(feedContent, prevFeedContent)) > 0, nil
}

func (a Atom) GetEntries(feedContent []byte) []entry {
	f := &feed{}

	// Parsing errors are handled like no entries exist
	_ = xml.Unmarshal(feedContent, f)

	return f.Entries
}

func (a Atom) GetNewEntries(feedContent []byte, prevFeedContent []byte) []entry {
	currEntries := a.GetEntries(feedContent)
	prevEntries := make(map[string]entry)
	for _, entry := range a.GetEntries(prevFeedContent) {
		prevEntries[entry.ID] = entry
	}

	changedEntries := make([]entry, 0)
	for _, entry := range currEntries {
		_, known := prevEntries[entry.ID]

		if !known {
			changedEntries = append(changedEntries, entry)
		}
	}

	return changedEntries
}

func (a Atom) GetChanges(feedContent []byte, prevFeedContent []byte) ([]interface{}, error) {
	newEntries := a.GetNewEntries(feedContent, prevFeedContent)
	changes := make([]interface{}, 0)

	for _, entry := range newEntries {
		changes = append(changes, entry)
	}

	return changes, nil
}
