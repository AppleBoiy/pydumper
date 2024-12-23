package scraping

import (
	"log"

	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func scrapingProcess(ch chan string, n int) {
	for {
		word := <-ch
		log.Printf("[Worker %d] : Get '%s' to work.", n, word)

		links, err := getURLFromWord(word)
		if err != nil {
			log.Printf("[%s] %s\n", word, err)
			continue
		}

		pages := []model.RawData{}
		for _, link := range links {
			page, err := getPage(link)
			if err != nil {
				log.Printf("[%s, URL: %s] %s\n", word, link, err)
				continue
			}
			pages = append(pages, model.RawData{Body: page, Link: link})
		}

		err = model.MD.Raw.Insert(pages, word)
		if err != nil {
			log.Printf("[%s] %s\n", word, err)
		}
	}
}
