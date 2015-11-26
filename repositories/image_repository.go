package repositories
import (
	"sportan/databases"
	"sportan/services"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2/bson"

)

type Image struct {
	ID string  `bson:"id,omitempty"`
	Content *string `bson:"content,omitempty"`
	BContent []byte `bson:"bcontent,omitempty"`
}

type ImageRepository struct {
	mongo *databases.MongoConfig
}

func NewImageRepository(mConfig *databases.MongoConfig) *ImageRepository {
	//Collection is name

	return &ImageRepository{
		mongo: mConfig,
	}
}


func (rep *ImageRepository) SaveImage(image *services.Image) (*services.Image, error){
	u, _ := uuid.NewV4()
	imgId := u.String()

	rep.mongo.Collection.Insert(
		&Image {
			ID : imgId,
			Content : image.Content,
			BContent : image.Bcontent,
		});

	image.ID = &imgId
	return image,nil

}

func (rep *ImageRepository) GetImageById(id string) (*services.Image) {
	var image Image
	rep.mongo.Collection.Find(bson.M{"id": id}).One(&image)

	return &services.Image {
		Content: image.Content,
		Bcontent: image.BContent,
		ID: &image.ID,
	}
}
