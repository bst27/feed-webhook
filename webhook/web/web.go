package web

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get(url string) *Web {
	return &Web{
		url: url,
	}
}

type Web struct {
	url string
}

func (w *Web) Receive(payload interface{}) error {
	content, err := json.MarshalIndent(payload, "", "   ")
	if err != nil {
		return err
	}

	resp, err := http.Post(w.url, "application/json", bytes.NewReader(content))
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
