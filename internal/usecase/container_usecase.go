package usecase

import (
	"database/sql"
	"golang-project-template/internal/domain"
	"golang-project-template/internal/repository"
)

type containerUseCase struct {
	repository domain.ContainerRepository
	mapper     *domain.ContainerMapper
}

func NewContainerUseCase(db *sql.DB, mapper *domain.ContainerMapper) domain.ContainerUseCase {

	return &containerUseCase{
		repository: repository.NewContainerPostgresRepository(db, &repository.ContainerPersistenceSqlMapper{}),
		mapper:     mapper,
	}
}

func (c containerUseCase) Create(request *domain.ContainerRequest) (int, error) {
	// check image exists
	container := c.mapper.MapToDomain(request)
	id, err := c.repository.Save(container)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c containerUseCase) GetById(id int) (*domain.ContainerDetailedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c containerUseCase) Filter(name string) (*[]domain.ContainerHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c containerUseCase) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (c containerUseCase) DeleteUnusedOnes() (float32, error) {
	//TODO implement me
	panic("implement me")
}
