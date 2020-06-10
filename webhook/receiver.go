package webhook

import (
	"github.com/bst27/feed-webhook/webhook/filesystem"
	"github.com/bst27/feed-webhook/webhook/stdout"
	"github.com/bst27/feed-webhook/webhook/web"
)

func Web(url string) Receiver {
	return web.Get(url)
}

func File(directory string) Receiver {
	return filesystem.Get(directory)
}

func Stdout() Receiver {
	return stdout.Get()
}

type Receiver interface {
	Receive(payload interface{}) error
}
