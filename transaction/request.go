package transaction

type RequestTransaction struct {
	DesribeComplaint string `validate:"required" json:"describe_complaint"`
	IdUser           int
	IdService        int
}
