package filesystem

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func Get(directory string) *Filesystem {
	return &Filesystem{
		directory: directory,
	}
}

type Filesystem struct {
	directory string
}

func (f *Filesystem) Receive(payload interface{}) error {
	file := fmt.Sprintf("%d-%s.json", time.Now().UnixNano(), uuid.New())
	path := filepath.Join(f.directory, file)

	content, err := json.MarshalIndent(payload, "", "   ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(f.directory, 0664)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, content, 0664)
}
