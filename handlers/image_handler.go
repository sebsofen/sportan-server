package handlers


import (

"sportan/repositories"
"sportan/services"

)

type ImageHandler struct {
	repo *repositories.ImageRepository
	userR *repositories.UserRepository
}

func NewImageHandler(repo *repositories.ImageRepository, userR *repositories.UserRepository) *ImageHandler {
	return &ImageHandler{
	repo:   repo,
	userR:  userR,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *ImageHandler) GetImageById(id string) (*services.Image, error) {
	image := ch.repo.GetImageById(id)

	return image,nil
}

func (ch *ImageHandler) CreateImage(token string, image *services.Image) (string,error) {
	userid, _ := ch.userR.GetUserIdFromToken(token)
	imgid := ""
	if ch.userR.IsAdmin(userid) {
		img, _ := ch.repo.SaveImage(image)
		imgid = *img.ID

	}
	return imgid, nil
}

