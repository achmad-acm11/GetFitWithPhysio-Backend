package galery

type GaleryResponse struct {
	Id         int    `json:"id"`
	Photo      string `json:"photo"`
	Caption    string `json:"caption"`
	SubCaption string `json:"sub_caption"`
}

func MapGaleryResponse(galery Galery) GaleryResponse {
	return GaleryResponse{
		Id:         galery.Id,
		Photo:      galery.Photo,
		Caption:    galery.Caption,
		SubCaption: galery.SubCaption,
	}
}
func MapGaleriesResponse(galeries []Galery) []GaleryResponse {
	var galeriesRes []GaleryResponse
	for _, v := range galeries {
		galeriesRes = append(galeriesRes, MapGaleryResponse(v))
	}

	return galeriesRes
}
