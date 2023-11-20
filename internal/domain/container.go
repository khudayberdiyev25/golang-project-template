package domain

import (
	"time"
)

type ContainerStatus int

const (
	ContainerUp = iota
	ContainerStopped
)

type Container struct {
	Id      int
	Name    string
	ImageId int
	Command string
	Created time.Time
	Status  ContainerStatus
}

type ContainerRequest struct {
	Name    string
	ImageId int
	Command string
}

type ContainerHeaderResponse struct {
	Id      int
	Name    string
	ImageId int
	Command string
	Created time.Time
	Status  ContainerStatus
}

type ContainerDetailedResponse struct {
	Id      int
	Name    string
	ImageId int
	Command string
	Created time.Time
	Status  ContainerStatus
}

type ContainerUseCase interface {
	Create(request *ContainerRequest) (int, error)
	GetById(id int) (*ContainerDetailedResponse, error)
	Filter(name string) (*[]ContainerHeaderResponse, error)
	DeleteById(id int) error
	DeleteUnusedOnes() (float32, error)
}

type ContainerRepository interface {
	Save(container *Container) (int, error)
	GetById(id int) (*ContainerDetailedResponse, error)
	Filter(name string) (*[]ContainerHeaderResponse, error)
	DeleteById(id int) error
	DeleteUnusedOnes() (float32, error)
}
