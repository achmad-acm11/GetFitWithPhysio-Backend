package exception

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(res http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(res, req, err) {
		return
	}

	if validationError(res, req, err) {
		return
	}
	internalServerError(res, req, err)
}
func validationError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		res.Header().Add("content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)

		response := helper.FormatResponse{
			Meta: helper.Meta{
				Message: "Incomplete Data",
				Status:  "Bad Request",
				Code:    http.StatusBadRequest,
			},
			Data: exception.Error(),
		}
		helper.WriteToResponsebody(res, response)

		return true
	} else {
		return false
	}
}
func notFoundError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		res.Header().Add("content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)

		response := helper.FormatResponse{
			Meta: helper.Meta{
				Message: "Data Not Found",
				Status:  "Not Found",
				Code:    http.StatusNotFound,
			},
			Data: exception.Error,
		}

		helper.WriteToResponsebody(res, response)

		return true
	} else {
		return false
	}
}

func internalServerError(res http.ResponseWriter, req *http.Request, err interface{}) {
	res.Header().Add("content-type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Internal Server Error",
			Status:  "Internal Server Error",
			Code:    http.StatusInternalServerError,
		},
		Data: err,
	}

	helper.WriteToResponsebody(res, response)

}
