package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type imagePostgresRepository struct {
	db sql.DB
}

func NewImagePostgresRepository(db sql.DB) domain.ImageRepository {
	return imagePostgresRepository{
		db: db,
	}
}

func (i imagePostgresRepository) Save(image domain.Image) (int, error) {
	var id int
	err := i.db.QueryRow(`insert into images(name) values ($1) returning id`, image.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (i imagePostgresRepository) FindAll(name string) ([]domain.Image, error) {
	rows, err := i.db.Query(`select * from images i where i.name = $1`, name)
	if err != nil {
		return nil, err
	}
	var images []domain.Image
	rows.Scan(&images)

	return images, nil
}

func (i imagePostgresRepository) GetByIdOrName(key string) domain.Image {
	var image domain.Image
	i.db.QueryRow(`select * from images i where i.id = $1 or i.name = $1`, key).Scan(&image)

	return image
}

func (i imagePostgresRepository) DeleteByIdOrName(key string) error {
	i.db.QueryRow(`delete from images where id = $1 or name = $1`, key)
	return nil
}

func (i imagePostgresRepository) DeleteAllUnused() float32 {
	i.db.QueryRow(`delete from images where not exists(select image_id from containers c where c.image_id = $1)`)
	return 0
}
