package testimonial

import (
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/user"
	"encoding/json"
)

type Testimonial struct {
	Id      int
	Id_user int
	Content string
	User    user.User       `gorm:"ForeignKey: Id_user"`
	Patient patient.Patient `gorm:"Foreignkey:Id_user;PrimaryKey:testimonial.Id_user;"`
}

type TmpTestimonial struct {
	Id         int
	Id_user    int
	Content    string
	Name       string
	Email      string
	Password   string
	Photo_user string
	Occupation string
	Nik        string
	Gender     string
	Phone      string
}

func MapTestimonial(result map[string]interface{}) Testimonial {
	byte, _ := json.Marshal(result)
	tmp := &TmpTestimonial{}
	json.Unmarshal(byte, tmp)

	return Testimonial{
		Id:      tmp.Id,
		Id_user: tmp.Id_user,
		Content: tmp.Content,
		User: user.User{
			Name:       tmp.Name,
			Email:      tmp.Email,
			Password:   tmp.Password,
			Photo_user: tmp.Photo_user,
		},
		Patient: patient.Patient{
			Occupation: tmp.Occupation,
			Nik:        tmp.Nik,
			Gender:     tmp.Gender,
			Phone:      tmp.Phone,
		},
	}
}
func MapTestimonials(result []map[string]interface{}) []Testimonial {
	var testimonials []Testimonial
	for _, v := range result {
		testimonials = append(testimonials, MapTestimonial(v))
	}
	return testimonials
}
