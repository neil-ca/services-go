package models

import (
	"fmt"
	"net/http"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

func (i *Item) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
