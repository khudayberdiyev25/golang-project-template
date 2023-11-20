package usecase

import (
	"database/sql"
	"fmt"
	"golang-project-template/internal/domain"
	"golang-project-template/internal/repository"
)

type imageUseCase struct {
	repository domain.ImageRepository
	mapper     domain.ImageMapper
}

func NewImageUseCase(db *sql.DB) domain.ImageUseCase {
	return &imageUseCase{
		repository: repository.NewImagePostgresRepository(db),
	}
}

func (i *imageUseCase) Create(request *domain.ImageRequest) (int, error) {
	id, err := i.repository.Save(&domain.Image{
		Name: request.Name,
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (i *imageUseCase) Filter(name string) (*[]domain.ImageHeaderResponse, error) {
	all, err := i.repository.FindAll(name)
	if err != nil {
		return nil, err
	}

	return i.mapper.MapToHeaderResponseSlice(all), nil
}

func (i *imageUseCase) DeleteByIdOrName(key string) error {
	err := i.repository.DeleteByIdOrName(key)
	if err != nil {
		return err
	}
	return nil
}

func (i *imageUseCase) DeleteAllUnusedOnes() (string, error) {
	reclaimedSpace, err := i.repository.DeleteAllUnused()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Total reclaimed space: %.2fB", reclaimedSpace), nil
}

func (i *imageUseCase) GetOne(key string) (*domain.ImageDetailedResponse, error) {
	image, err := i.repository.GetByIdOrName(key)
	if err != nil {
		return &domain.ImageDetailedResponse{}, err
	}
	return i.mapper.MapToDetailedResponse(image), nil
}
