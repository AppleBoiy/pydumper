package scraping

import (
	"log"
	"strings"

	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func scrapingProcess(ch chan string, n int, stop chan bool) {
	for {
		word := <-ch
		if word == "." {
			stop <- true
			break
		}
		log.Printf("[Worker %d] : Get '%s' to work.", n, word)

		links, err := getURLFromWord(word)
		if err != nil {
			if strings.Contains(err.Error(), "429") {
				stop <- true
				break
			}
			log.Printf("[%s] %s\n", word, err)
			continue
		}

		pages := []model.RawData{}
		for _, link := range links {
			// get free space
			// f_space, err := util.GetFreeSpace()
			// if err != nil {
			// 	log.Printf("[%s, URL: %s] %s\n", word, link, err)
			// }

			// check free space
			// if f_space < 1e9 {
			// 	log.Printf("[%s, URL: %s] %s\n", word, link, "not enough space (1GB left).")
			// 	return
			// }

			// get page
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
