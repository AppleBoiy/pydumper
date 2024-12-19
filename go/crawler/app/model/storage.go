package model

import (
	"database/sql"
)

type rawData struct {
	db *sql.DB
}

type RawData struct {
	Body string `json:"body"`
	Link string `json:"link"`
}

func (r *rawData) initialize(db *sql.DB) {
	r.db = db
}

func (r *rawData) Insert(raw []RawData) error {
	sql_cmd := `INSERT INTO raw_data
	()
	`
	return nil
}
