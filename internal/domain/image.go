package domain

type Image struct {
	Id   int
	Name string
}

type ImageRequest struct {
	Name string
}

type ImageDetailedResponse struct {
	Id   int
	Name string
}

type ImageHeaderResponse struct {
	Id   int
	Name string
}

type ImageUseCase interface {

	// Create docker build
	Create(request ImageRequest) (int, error)

	// Filter docker images
	Filter(name string) ([]ImageHeaderResponse, error)

	// DeleteByIdOrName docker rmi
	DeleteByIdOrName(key string) error

	// DeleteAllUnusedOnes docker image prune
	// returns total reclaimed space
	DeleteAllUnusedOnes() string

	// GetOne docker image inspect
	GetOne(key string) (ImageDetailedResponse, error)
}

type ImageRepository interface {
	Save(image Image) (int, error)
	FindAll(name string) ([]Image, error)
	GetByIdOrName(key string) Image
	DeleteByIdOrName(key string) error
	DeleteAllUnused() float32
}
