package stdout

import (
	"encoding/json"
	"fmt"
)

func Get() *Stdout {
	return &Stdout{}
}

type Stdout struct {
}

func (s *Stdout) Receive(payload interface{}) error {
	content, err := json.MarshalIndent(payload, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(content))
	return nil
}
