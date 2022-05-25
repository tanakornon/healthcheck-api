package healthcheck

import (
	"backend/src/utils"
	"encoding/json"
	"net/http"
	"strings"
)

type IController interface {
	FindSiteStatus(w http.ResponseWriter, r *http.Request)
}

type healthCheckController struct {
	healthCheckService
}

func NewController() healthCheckController {
	return healthCheckController{healthCheckService{}}
}

func (controller healthCheckController) FindSiteStatus(w http.ResponseWriter, r *http.Request) {
	raw := utils.ReceiveFileContent(r)

	var siteList []string

	for _, site := range strings.Split(raw, ",") {
		url := strings.Trim(site, " \n")
		if url != "" {
			siteList = append(siteList, url)
		}
	}

	result := controller.PingSites(siteList)
	json.NewEncoder(w).Encode(result)
}
