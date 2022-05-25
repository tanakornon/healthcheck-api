package apis

import (
	"backend/src/apis/common"
	"backend/src/apis/healthcheck"
)

type Controllers struct {
	Common      common.IController
	HealthCheck healthcheck.IController
}

func NewControllers() Controllers {
	return Controllers{
		Common:      common.NewController(),
		HealthCheck: healthcheck.NewController(),
	}
}
