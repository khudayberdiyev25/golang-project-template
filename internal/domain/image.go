package domain

type Image struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ImageRequest struct {
	Name string
}

type ImageDetailedResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ImageHeaderResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ImageUseCase interface {

	// Create docker build
	Create(request *ImageRequest) (int, error)

	// Filter docker images
	Filter(name string) (*[]ImageHeaderResponse, error)

	// DeleteByIdOrName docker rmi
	DeleteByIdOrName(key string) error

	// DeleteAllUnusedOnes docker image prune
	// returns total reclaimed space
	DeleteAllUnusedOnes() (string, error)

	// GetOne docker image inspect
	GetOne(key string) (*ImageDetailedResponse, error)
}

type ImageRepository interface {
	Save(image *Image) (int, error)
	FindAll(name string) (*[]Image, error)
	GetByIdOrName(key string) (*Image, error)
	DeleteByIdOrName(key string) error
	DeleteAllUnused() (float32, error)
}
