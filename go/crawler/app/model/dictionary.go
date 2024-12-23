package model

import (
	"database/sql"
	"log"
	"strings"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/util"
)

type dictionary struct {
	db *sql.DB
}

func (d *dictionary) Initialize(db *sql.DB) {
	d.db = db
}

func (d *dictionary) GetAllDictionary() []string {
	result := []string{}

	sql_cmd := `SELECT word FROM words;`

	rows, err := d.db.Query(sql_cmd)
	if err != nil {
		util.ErrLog("failed to get all words:", err)
	}

	defer rows.Close()

	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, strings.TrimSpace(word))
	}

	return result
}
