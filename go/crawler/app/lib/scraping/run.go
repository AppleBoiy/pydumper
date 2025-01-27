package scraping

import (
	"log"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"
)

func RunWorker(ch chan string, end chan bool) {
	t_num := config.Config.Thread
	if t_num < 1 {
		t_num = 1
	}
	log.Println("Running Workers.")

	p_stop := make(chan bool, 1)
	for i := 0; i < t_num; i++ {
		go scrapingProcess(ch, i+1, p_stop)
	}

	for i := 0; i < t_num; i++ {
		<-p_stop
	}
	end <- true
}
