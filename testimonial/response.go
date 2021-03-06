package testimonial

type TestimonialResponse struct {
	Id         int    `json:"id"`
	Id_user    int    `json:"id_user"`
	Photo_user string `json:"photo_user"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Content    string `json:"content"`
	Occupation string `json:"occupation"`
}

func MapTestimonialResponse(testimonial Testimonial) TestimonialResponse {
	return TestimonialResponse{
		Id:         testimonial.Id,
		Id_user:    testimonial.Id_user,
		Photo_user: testimonial.User.Photo_user,
		Content:    testimonial.Content,
		Name:       testimonial.User.Name,
		Email:      testimonial.User.Email,
		Occupation: testimonial.Patient.Occupation,
	}
}
func MapTestimonialsResponse(testimonials []Testimonial) []TestimonialResponse {
	var testimonialsRes []TestimonialResponse
	for _, v := range testimonials {
		testimonialsRes = append(testimonialsRes, MapTestimonialResponse(v))
	}
	return testimonialsRes
}
