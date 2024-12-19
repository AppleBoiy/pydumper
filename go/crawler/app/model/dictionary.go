package model

import "database/sql"

type dictionary struct {
	db *sql.DB
}

func (d *dictionary) initialize(db *sql.DB) {
	d.db = db
}

func (d *dictionary) GetAllDictionary() []string {
	return []string{}
}
