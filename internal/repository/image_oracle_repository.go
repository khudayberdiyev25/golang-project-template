package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type imageOracleRepository struct {
	db sql.DB
}

func NewImageOracleRepository(db sql.DB) domain.ImageRepository {
	return imageOracleRepository{
		db: db,
	}
}

func (i imageOracleRepository) Save(image domain.Image) (int, error) {
	panic("implement me")
}

func (i imageOracleRepository) FindAll(name string) ([]domain.Image, error) {
	panic("implement me")
}

func (i imageOracleRepository) GetByIdOrName(key string) domain.Image {
	panic("implement me")
}

func (i imageOracleRepository) DeleteByIdOrName(key string) error {
	panic("implement me")

}

func (i imageOracleRepository) DeleteAllUnused() float32 {
	panic("implement me")
}
