package transaction

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type controllerTransaction struct {
	service ServiceTransaction
}

func NewControllerTransaction(service ServiceTransaction) *controllerTransaction {
	return &controllerTransaction{
		service: service,
	}
}

// Create Trasaction
func (c *controllerTransaction) CreateTransactionController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceId, err := strconv.Atoi(params.ByName("serviceId"))
	helper.HandleError(err)

	requestTransaction := RequestTransaction{}
	helper.ReadFromRequestBody(req, &requestTransaction)

	requestTransaction.IdService = serviceId
	requestTransaction.IdUser = 10

	data := c.service.CreateService(req.Context(), requestTransaction)

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Transaction Success",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}

	helper.WriteToResponsebody(res, response)
}

func (c *controllerTransaction) GetAllTransactionController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	data := c.service.GetAllService(req.Context())

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Get Success",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}

	helper.WriteToResponsebody(res, response)
}
