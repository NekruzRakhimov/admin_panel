package models

import "time"

type Dictionary struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	LastUpdate time.Time `json:"last_update"`
	Author     string    `json:"author"`
}

type DictionaryValue struct {
	ID           int    `json:"id"`
	Value        string `json:"value"`
	DictionaryID int    `json:"dictionary_id"`
}
