package healthcheck

type Health string

const (
	Up   Health = "up"
	Down        = "down"
)

type SiteStatus struct {
	Url    string `json:"url"`
	Health Health `json:"health"`
}

type Results struct {
	SitesStatus []SiteStatus `json:"sites_status"`
	Elapsed     int64        `json:"elapsed"`
}
