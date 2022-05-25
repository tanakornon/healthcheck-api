package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockTestServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	service = healthCheckService{}
)

func TestUpSite(t *testing.T) {
	ts := mockTestServer
	defer ts.Close()

	ch := make(chan SiteStatus)

	go service.ping(ts.URL, ch)

	result := <-ch
	if result.Health != Up {
		t.Fail()
	}
}

func TestDownSite(t *testing.T) {
	ch := make(chan SiteStatus)

	go service.ping("test", ch)

	result := <-ch
	if result.Health != Down {
		t.Fail()
	}
}

func TestPingSites(t *testing.T) {
	ts := mockTestServer
	defer ts.Close()

	urls := []string{ts.URL, "random_url"}

	result := service.PingSites(urls)

	if len(result.SitesStatus) != 2 {
		t.Fail()
	}
}
