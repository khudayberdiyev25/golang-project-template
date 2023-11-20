package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-project-template/internal/domain"
)

type imagePostgresRepository struct {
	db *sql.DB
}

func NewImagePostgresRepository(db *sql.DB) domain.ImageRepository {
	return &imagePostgresRepository{
		db: db,
	}
}

func (i *imagePostgresRepository) Save(image *domain.Image) (int, error) {
	var id int
	err := i.db.QueryRow(`insert into images(name) values ($1) returning id`, image.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (i *imagePostgresRepository) FindAll(name string) (*[]domain.Image, error) {
	baseQueryStr := `select * from images`
	rows, err := i.db.Query(baseQueryStr)
	if err != nil {
		return nil, err
	}
	var images []domain.Image
	for rows.Next() {
		image := domain.Image{}
		err := rows.Scan(&image.Id, &image.Name)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return &images, nil
}

func (i *imagePostgresRepository) GetByIdOrName(key string) (*domain.Image, error) {
	image := domain.Image{}
	err := i.db.QueryRow(`select * from images i where i.id = $1 or i.name = $2`, key, key).Scan(&image.Id, &image.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Image{}, &domain.ErrImageNotFound{
				Err:    fmt.Sprintf("image not found with %v key", key),
				Status: 400,
			}
		} else {
			return &domain.Image{}, err
		}
	}
	return &image, nil
}

func (i *imagePostgresRepository) DeleteByIdOrName(key string) error {
	_, err := i.db.Exec(`delete from images where id = $1 or name = $2`, key, key)
	if err != nil {
		return err
	}
	return nil
}

func (i *imagePostgresRepository) DeleteAllUnused() (float32, error) {
	_, err := i.db.Exec(`delete from images`)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
