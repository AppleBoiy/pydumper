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
	end := make(chan bool, 1)
	go scraping.RunWorker(ch, end)

	// Main Procress
	dicts := model.MD.Dict.GetAllDictionary()
	if config.Config.MAXSearch != 0 && len(dicts) > config.Config.MAXSearch {
		dicts = dicts[:config.Config.MAXSearch]
	}

	for _, dict := range dicts {
		log.Printf("Pushing Word: %s\n", dict)
		ch <- dict
	}

	log.Println("Finish Process!!")
	<-end
}
