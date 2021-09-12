package client

import "fmt"

type SearchEngine struct {
	Name      string
	SearchUrl string
}

func (s *SearchEngine) FillUrl(request string) {
	s.SearchUrl = fmt.Sprintf(s.SearchUrl, request)
}

func GetSearchEngines() []SearchEngine {
	return []SearchEngine{
		{"google", "https://www.google.com/search?q=%s"},
		{"bing", "https://www.bing.com/search?q=%s"},
		{"yahoo", "https://search.yahoo.com/search?p=%s"},
		{"yandex", "https://yandex.com/search/?text=%s"},
		{"mailru", "https://go.mail.ru/search?q=%s"},
		{"ramblerru", "https://nova.rambler.ru/search?query=%s"},
		{"duckduckgo", "https://duckduckgo.com/?q=%s"},
		{"fireball", "https://fireball.de/search?q=%s"},
		{"ask", "https://www.ask.com/web?q=%s"},
		{"metacrawler", "https://www.metacrawler.com/serp?q=%s"},
	}
}
