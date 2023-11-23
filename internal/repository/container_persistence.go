package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type ContainerSQL struct {
	Id      int
	Name    sql.NullString
	ImageId int
	Command sql.NullString
	Created sql.NullTime
	Status  int
}

type ContainerPersistenceSqlMapper struct {
}

func (c *ContainerPersistenceSqlMapper) MapToDomainFromSql(source *ContainerSQL) *domain.Container {

	return &domain.Container{
		Id:      source.Id,
		Name:    source.Name.String,
		ImageId: source.ImageId,
		Command: source.Command.String,
		Created: source.Created.Time,
		Status:  domain.ContainerStatus(source.Status),
	}
}
