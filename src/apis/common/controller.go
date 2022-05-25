package common

import (
	"encoding/json"
	"net/http"
)

type IController interface {
	Pong(w http.ResponseWriter, r *http.Request)
}

type commonController struct{}

func NewController() commonController {
	return commonController{}
}

func (controller commonController) Pong(w http.ResponseWriter, r *http.Request) {
	data := struct{ Message string }{"Pong"}
	json.NewEncoder(w).Encode(data)
}
