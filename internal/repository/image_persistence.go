package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type ImageSQL struct {
	Id      int
	Name    sql.NullString
	Created sql.NullTime
	RepoTag sql.NullString
}

type ImagePersistenceSqlMapper struct {
}

func (m *ImagePersistenceSqlMapper) MapToDomainFromSql(source *ImageSQL) *domain.Image {
	return &domain.Image{
		Id:      source.Id,
		Name:    source.Name.String,
		Created: source.Created.Time,
		RepoTag: source.RepoTag.String,
	}
}

func (m *ImagePersistenceSqlMapper) MapToDomainSliceFromSql(sourceSlice *[]ImageSQL) *[]domain.Image {
	targetSlice := []domain.Image{}
	for _, imageSQL := range *sourceSlice {
		targetSlice = append(targetSlice, *m.MapToDomainFromSql(&imageSQL))
	}
	return &targetSlice
}
