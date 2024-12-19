package scraping

import "github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"

func RunWorker(ch chan string) {
	t_num := config.Config.Thread
	if t_num < 1 {
		t_num = 1
	}

	for i := 0; i < t_num; i++ {
		go scrapingProcess(ch)
	}
}
