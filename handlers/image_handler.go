package handlers


import (

"sportan/repositories"
"sportan/services"

)

type ImageHandler struct {
repo *repositories.ImageRepository
}

func NewImageHandler(repo *repositories.ImageRepository) *ImageHandler {
	return &ImageHandler{
	repo:   repo,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *ImageHandler) GetImageById(id string) (*services.Image, error) {
	image := ch.repo.GetImageById(id)

	return image,nil
}

