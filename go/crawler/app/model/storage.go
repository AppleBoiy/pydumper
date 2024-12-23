package model

import (
	"database/sql"
	"fmt"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/util"
)

type rawData struct {
	db *sql.DB
}

type RawData struct {
	Body string `json:"body"`
	Link string `json:"link"`
}

func (r *rawData) Initialize(db *sql.DB) {
	r.db = db
}

func (r *rawData) Insert(raw []RawData, keyword string) error {
	if len(raw) == 0 {
		return nil
	}

	sql_cmd := `INSERT INTO raw_data
	(keyword, site, raw)
	VALUES %s;
	`

	sql_cmd = fmt.Sprintf(sql_cmd, util.SqlGenGroup(len(raw), 3, 1))

	arg := []any{}
	for _, it := range raw {
		arg = append(arg, keyword, it.Link, it.Body)
	}

	_, err := r.db.Exec(sql_cmd, arg...)
	return err
}
