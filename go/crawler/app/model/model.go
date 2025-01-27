package model

import (
	"database/sql"
	"fmt"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"
	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/util"
	_ "github.com/lib/pq"
)

type Model struct {
	Dict *dictionary
	Raw  *rawData
}

func (m *Model) Initialize() {
	connect_info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config.DBHOST, config.Config.DBPORT, config.Config.DBUSER, config.Config.DBPASS, config.Config.DBNAME)

	db, err := sql.Open("postgres", connect_info)
	if err != nil {
		util.ErrLog("failed to connect database:", err)
	}
	m.Dict = &dictionary{}
	m.Raw = &rawData{}
	m.Dict.Initialize(db)
	m.Raw.Initialize(db)
}

var MD Model = Model{}
