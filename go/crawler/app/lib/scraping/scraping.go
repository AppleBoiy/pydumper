package scraping

import (
	"log"
	"strings"

	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func scrapingProcess(n int, stop chan bool) {
	for {
		id, word, err := model.MD.Dict.GetOneWord()
		if err != nil {
			log.Println("[Get Word Error]: ", err)
			continue
		}

		links, err := getURLFromWord(word)
		if err != nil {
			model.MD.Dict.OnWordError(id)
			if strings.Contains(err.Error(), "429") {
				log.Printf("Worker %d has done.", n)
				stop <- true
				break
			}
			log.Printf("[%s] %s\n", word, err)
			continue
		}

		pages := []model.RawData{}
		for _, link := range links {
			// get page
			page, err := getPage(link)
			if err != nil {
				// log.Printf("[%s] %s\n", word, link, err)
				continue
			}
			pages = append(pages, model.RawData{Body: page, Link: link})
		}

		err = model.MD.Raw.Insert(pages, word)
		if err != nil {
			model.MD.Dict.OnWordError(id)
			log.Printf("[%s] %s\n", word, err)
			continue
		}

		err = model.MD.Dict.OnWordDone(id)
		if err != nil {
			log.Printf("[%s] %s\n", word, err)
			continue
		}

		log.Printf("\"%s\" have done.\n", word)
	}
}
