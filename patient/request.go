package patient

type RegisterRequest struct {
	Name       string `validate:"required" json:"name"`
	Gender     string `validate:"required" json:"gender"`
	Phone      string `validate:"required" json:"phone"`
	Address    string `validate:"required" json:"address"`
	Nik        string `validate:"required" json:"nik"`
	Birth_date string `json:"birthdate"`
	Email      string `validate:"required" json:"email"`
	Password   string `validate:"required" json:"password"`
	Role       int
}
type CreatePatientRequest struct {
	Name       string `validate:"required" json:"name"`
	Gender     string `validate:"required" json:"gender"`
	Phone      string `validate:"required" json:"phone"`
	Address    string `validate:"required" json:"address"`
	Nik        string `validate:"required" json:"nik"`
	Birth_date string `json:"birthdate"`
	Email      string `validate:"required" json:"email"`
	Password   string `validate:"required" json:"password"`
	Role       int
}
