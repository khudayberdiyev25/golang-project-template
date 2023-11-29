package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type containerPostgresRepository struct {
	db     *sql.DB
	mapper *ContainerPersistenceSqlMapper
}

func NewContainerPostgresRepository(db *sql.DB, mapper *ContainerPersistenceSqlMapper) domain.ContainerRepository {

	return &containerPostgresRepository{
		db:     db,
		mapper: mapper,
	}
}

func (c containerPostgresRepository) Save(container *domain.Container) (int, error) {
	var id int
	err := c.db.QueryRow(`insert into containers(name, image_id, command, created, status) values ($1, $2, $3, $4, $5) returning id`,
		container.Name, container.ImageId, container.Command, container.Created, container.Status).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c containerPostgresRepository) GetById(id int) (*domain.ContainerDetailedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c containerPostgresRepository) Filter(name string) (*[]domain.ContainerHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c containerPostgresRepository) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (c containerPostgresRepository) DeleteUnusedOnes() (float32, error) {
	//TODO implement me
	panic("implement me")
}
