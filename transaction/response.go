package transaction

type TransactionResponse struct {
	Id               int    `json:"id"`
	IdUser           int    `json:"id_user"`
	IdService        int    `json:"id_service"`
	Name             string `json:"name"`
	ServiceName      string `json:"service_name"`
	DesribeComplaint string `json:"describe_complaint"`
	Code             string `json:"code"`
	Amount           int    `json:"amount"`
	Status           string `json:"status"`
}

func MapTransactionResponse(transaction Transaction) TransactionResponse {
	return TransactionResponse{
		Id:               transaction.Id,
		IdUser:           transaction.Id_user,
		IdService:        transaction.Id_service,
		Name:             transaction.User.Name,
		ServiceName:      transaction.Service.Service_name,
		DesribeComplaint: transaction.DesribeComplaint,
		Code:             transaction.Code,
		Amount:           transaction.Amount,
		Status:           transaction.Status,
	}
}

func MapTransactionsResponse(transactions []Transaction) []TransactionResponse {
	transactionsRes := []TransactionResponse{}

	for _, v := range transactions {
		transactionsRes = append(transactionsRes, MapTransactionResponse(v))
	}

	return transactionsRes
}
