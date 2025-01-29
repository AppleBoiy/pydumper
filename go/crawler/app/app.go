package app

import (
	"log"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/scraping"
	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func Start() {
	// Reprepare
	err := model.MD.Dict.Reprepare()
	if err != nil {
		log.Println("[Reprepare Error]:", err)
		return
	}

	// Prepare Worker
	scraping.RunWorker()
}
