package testimonial

import (
	"GetfitWithPhysio-backend/user"
)

type Testimonial struct {
	Id      int
	Id_user int
	Content string
	User    user.User `gorm:"ForeignKey: Id_user"`
}
