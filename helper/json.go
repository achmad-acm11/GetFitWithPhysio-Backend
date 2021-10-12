package helper

import (
	"encoding/json"
	"net/http"
)

type FormatResponse struct {
	Meta Meta
	Data interface{}
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func WriteToResponsebody(wr http.ResponseWriter, response interface{}) {
	wr.Header().Add("content-type", "application/json")
	err := json.NewEncoder(wr).Encode(response)
	HandleError(err)
}
