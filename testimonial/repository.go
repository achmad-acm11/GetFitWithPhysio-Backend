package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryTestimonial interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Testimnoial
}

type repositoryTestimonial struct {
}

func NewRepositoryTestimonial() *repositoryTestimonial {
	return &repositoryTestimonial{}
}

// SQL Query Get All Testimonial
func (r *repositoryTestimonial) GetAll(ctx context.Context, tx *sql.Tx) []Testimnoial {
	query := "SELECT * FROM testimonials"

	data, err := tx.QueryContext(ctx, query)
	helper.HandleError(err)
	defer data.Close()

	var testimonials []Testimnoial
	for data.Next() {
		testimonial := Testimnoial{}

		err := data.Scan(&testimonial.Id, &testimonial.Id_user, &testimonial.Content)
		helper.HandleError(err)

		testimonials = append(testimonials, testimonial)
	}
	return testimonials
}
