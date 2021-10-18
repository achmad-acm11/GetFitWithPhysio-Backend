package galery

type CreateGaleryRequest struct {
	Caption    string `validate:"required" form:"caption"`
	SubCaption string `validate:"required" form:"sub_caption"`
}
