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

func (d *dictionary) GetOneWord() (int, string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return 0, "", err
	}

	sql_cmd := `
	SELECT id, word
	FROM mix_words
	WHERE status = $1
	LIMIT 1;
	`

	row := tx.QueryRow(sql_cmd, "a")
	var id int
	var word string
	err = row.Scan(&id, &word)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}

	sql_cmd = `
	UPDATE mix_words
	SET status = $1
	WHERE id = $2;
	`

	_, err = tx.Exec(sql_cmd, "p", id)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}

	return id, strings.Trim(word, "\n "), nil
}

func (d *dictionary) OnWordDone(id int) error {
	sql_cmd := `
	UPDATE mix_words
	SET status = $1
	WHERE id = $2;
	`

	_, err := d.db.Exec(sql_cmd, "d", id)
	return err
}

func (d *dictionary) OnWordError(id int) error {
	sql_cmd := `
	UPDATE mix_words
	SET status = $1
	WHERE id = $2;
	`
	_, err := d.db.Exec(sql_cmd, "a", id)
	return err
}

func (d *dictionary) Reprepare() error {
	sql_cmd := `
	UPDATE mix_words
	SET status = $1
	WHERE status = $2;
	`

	_, err := d.db.Exec(sql_cmd, "a", "p")
	return err
}

func (d *dictionary) GetAllDictionary() []string {
	result := []string{}

	sql_cmd := `SELECT word FROM mix_words;`

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
