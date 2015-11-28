package handlers


import (

	"sportan/repositories"
	"sportan/services"
	"github.com/nfnt/resize"
	"image/jpeg"

	"bytes"
	"bufio"
	"fmt"
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
	fmt.Println("Delivering Image " + id)
	image := ch.repo.GetImageById(id)

	return image,nil
}

func (ch *ImageHandler) GetThumbnailByImageId(id string) (*services.Image, error) {
	image, err  := ch.GetImageById(id);

	if(err != nil && image.Bcontent != nil) {
		//try to resize image bcontent
		//TODO assumption: this is allways a jpeg file!
		img, _ := jpeg.Decode(bytes.NewReader(image.Bcontent))
		m := resize.Resize(50, 50, img, resize.NearestNeighbor)
		var b bytes.Buffer
		writer := bufio.NewWriter(&b)
		jpeg.Encode(writer,m,nil)
		writer.Write(image.Bcontent)
		return image, err
	}else{
		return nil, err
	}


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

