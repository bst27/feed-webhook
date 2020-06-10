package reader

import (
	"crypto/md5"
	"github.com/bst27/feed-webhook/reader/atom"
)

type Reader interface {
	HasChanged(feedContent []byte, prevFeedContent []byte) (bool, error)
	GetChanges(feedContent []byte, prevFeedContent []byte) ([]interface{}, error)
}

type reader struct {
}

func GetAtom() Reader {
	return atom.Get()
}

func Get() Reader {
	return &reader{}
}

func (r reader) HasChanged(feedContent []byte, prevFeedContent []byte) (bool, error) {
	return md5.Sum(feedContent) != md5.Sum(prevFeedContent), nil
}

func (r reader) GetChanges(feedContent []byte, prevFeedContent []byte) ([]interface{}, error) {
	payload := make([]interface{}, 0)

	payload = append(payload, string(feedContent), string(prevFeedContent))

	return payload, nil
}
