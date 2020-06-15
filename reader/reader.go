package reader

import (
	"github.com/bst27/feed-webhook/reader/atom"
	"github.com/bst27/feed-webhook/reader/rss"
)

type Reader interface {
	HasChanged(feedContent []byte, prevFeedContent []byte) (bool, error)
	GetChanges(feedContent []byte, prevFeedContent []byte) ([]interface{}, error)
}

func GetAtom() Reader {
	return atom.Get()
}

func GetRSS() Reader {
	return rss.Get()
}
