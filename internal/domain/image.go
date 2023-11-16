package domain

type Image struct {
	Id   int64
	Name string
}

type ImageRequest struct {
	Name string
}

type ImageDetailedResponse struct {
	Name string
}

type ImageHeaderResponse struct {
	Name string
}

type ImageUseCase interface {

	// Create docker build
	Create(request ImageRequest) (int64, error)

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
