package healthcheck

import (
	"backend/src/utils"
	"net/http"
	"time"
)

type healthCheckService struct{}

func (service *healthCheckService) ping(url string, ch chan<- SiteStatus) {
	_, err := http.Get(url)
	health := Up

	if err != nil {
		health = Down
	}

	ch <- SiteStatus{url, health}
}

func (service *healthCheckService) PingSites(urls []string) Results {
	start := time.Now()

	var result Results

	ch := make(chan SiteStatus)

	for _, url := range urls {
		url = utils.ParseUrl(url)
		go service.ping(url, ch)
	}

	for range urls {
		result.SitesStatus = append(result.SitesStatus, <-ch)
	}

	result.Elapsed = time.Since(start).Milliseconds()

	return result
}
