package testimonial

type RequestTestimonial struct {
	Content string `validate:"required" json:"content"`
	Id_user int
}
