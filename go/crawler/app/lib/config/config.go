package config

import (
	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/util"
	"gopkg.in/ini.v1"
)

type config struct {
	APIKey         string `ini:"api_key"`
	SearchEngineID string `ini:"search_engine_id"`
	Thread         int    `ini:"thread"`
	WebMax         int    `ini:"web_number"`
	MAXSearch      int    `ini:"max_search"`
	DBHOST         string `ini:"db_host"`
	DBUSER         string `ini:"db_user"`
	DBPASS         string `ini:"db_pass"`
	DBNAME         string `ini:"db_name"`
	DBPORT         int    `ini:"db_port"`
}

var Config config

func LoadConfig() {
	cfg, err := ini.Load(util.GetPath("config.ini"))
	if err != nil {
		util.ErrLog("Load Config Error :", err)
	}

	if err := cfg.Section("default").MapTo(&Config); err != nil {
		util.ErrLog("Map Config Error :", err)
	}
}
