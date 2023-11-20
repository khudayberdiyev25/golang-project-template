package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-project-template/internal/domain"
	"log"
	"time"
)

type imagePostgresRepository struct {
	imagePersistenceSqlMapper *ImagePersistenceSqlMapper
	db                        *sql.DB
}

func NewImagePostgresRepository(db *sql.DB, imagePersistenceSqlMapper *ImagePersistenceSqlMapper) domain.ImageRepository {
	return &imagePostgresRepository{
		db:                        db,
		imagePersistenceSqlMapper: imagePersistenceSqlMapper,
	}
}

func (i *imagePostgresRepository) Save(image *domain.Image) (int, error) {
	var id int
	err := i.db.QueryRow(`insert into images(name, repo_tag, created) values ($1, $2, $3) returning id`, image.Name, image.RepoTag, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (i *imagePostgresRepository) FindAll(name string) (*[]domain.Image, error) {
	baseQueryStr := `select i.id, i.name, i.repo_tag, i.created from images i`
	rows, err := i.db.Query(baseQueryStr)
	if err != nil {
		return nil, err
	}
	var images []ImageSQL
	for rows.Next() {
		image := ImageSQL{}
		err := rows.Scan(&image.Id, &image.Name, &image.RepoTag, &image.Created)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return i.imagePersistenceSqlMapper.MapToDomainSliceFromSql(&images), nil
}

func (i *imagePostgresRepository) GetByIdOrName(key string) (*domain.Image, error) {
	image := ImageSQL{}
	err := i.db.QueryRow(`select i.id, i.name, i.repo_tag, i.created from images i where i.id = $1 or i.name = $2`, key, key).Scan(&image.Id, &image.Name, &image.RepoTag, &image.Created)
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
	log.Print(image.RepoTag)
	return i.imagePersistenceSqlMapper.MapToDomainFromSql(&image), nil
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
