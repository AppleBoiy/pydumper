package app

import (
	"log"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"
	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/scraping"
	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func Start() {
	// Prepare Worker
	ch := make(chan string, 1)
	scraping.RunWorker(ch)

	// Main Procress
	dicts := model.MD.Dict.GetAllDictionary()
	if len(dicts) > config.Config.MAXSearch {
		dicts = dicts[:config.Config.MAXSearch]
	}

	for _, dict := range dicts {
		log.Printf("Pushing Word: %s\n", dict)
		ch <- dict
	}
	log.Println("Finish Procress!!")
	for {
	}
}
