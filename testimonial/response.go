package testimonial

type TestimonialResponse struct {
	Id      int    `json:"id"`
	Id_user int    `json:"id_user"`
	Content string `json:"content"`
}

func MapTestimonialResponse(testimonial Testimnoial) TestimonialResponse {
	return TestimonialResponse{
		Id:      testimonial.Id,
		Id_user: testimonial.Id_user,
		Content: testimonial.Content,
	}
}
func MapTestimonialsResponse(testimonials []Testimnoial) []TestimonialResponse {
	var testimonialsRes []TestimonialResponse
	for _, v := range testimonials {
		testimonialsRes = append(testimonialsRes, MapTestimonialResponse(v))
	}
	return testimonialsRes
}
