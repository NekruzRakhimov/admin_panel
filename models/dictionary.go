package models

import "time"

type Dictionary struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	LastUpdate time.Time `json:"last_update"`
	Author     string    `json:"author"`
	HasFile bool `json:"has_file"`
}

type DictionaryValue struct {
	ID           int    `json:"id"`
	Value        string `json:"value"`
	File         string `json:"file"`
	HasFile bool `json:"has_file"`
	DictionaryID int    `json:"dictionary_id"`

}
