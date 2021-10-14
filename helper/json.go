package helper

import (
	"encoding/json"
	"net/http"
)

type FormatResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func ReadFromRequestBody(req *http.Request, result interface{}) {
	err := json.NewDecoder(req.Body).Decode(result)
	HandleError(err)
}
func WriteToResponsebody(wr http.ResponseWriter, response interface{}) {
	wr.Header().Add("content-type", "application/json")
	err := json.NewEncoder(wr).Encode(response)
	HandleError(err)
}
