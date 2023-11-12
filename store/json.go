package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JSONFileStore struct {
	FilePath string
}

func NewJSONFileStore(filePath string) *JSONFileStore {
	return &JSONFileStore{FilePath: filePath}
}

func (store *JSONFileStore) Create(item interface{}) error {
	items := make([]interface{}, 0)
	err := store.FindAll(&items)
	if err != nil {
		return err
	}
	items = append(items, item)
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return ioutil.WriteFile(store.FilePath, data, 0644)
}

func (store *JSONFileStore) FindAll(items interface{}) error {
	data, err := ioutil.ReadFile(store.FilePath)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, items)
}
