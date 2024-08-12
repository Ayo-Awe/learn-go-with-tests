package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	urlStatuses := make(map[string]bool)

	type result struct {
		url    string
		status bool
	}

	resultsChan := make(chan result)
	for _, url := range urls {
		go func() {
			status := wc(url)
			resultsChan <- result{url, status}
		}()
	}

	for range len(urls) {
		res := <-resultsChan
		urlStatuses[res.url] = res.status
	}

	return urlStatuses
}
