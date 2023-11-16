package usecase

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

type imageUseCase struct {
	db sql.DB
}

func NewImageUseCase(db sql.DB) domain.ImageUseCase {
	return &imageUseCase{
		db: db,
	}
}

func (i *imageUseCase) Create(request domain.ImageRequest) (int64, error) {
	var id int64
	err := i.db.QueryRow(`insert into images(name) values ($1) returning id`, request.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (i *imageUseCase) Filter(name string) ([]domain.ImageHeaderResponse, error) {
	rows, err := i.db.Query(`select * from images i where i.name = $1`, name)
	if err != nil {
		return nil, err
	}
	var images []domain.ImageHeaderResponse
	rows.Scan(&images)

	return images, nil
}

func (i *imageUseCase) DeleteByIdOrName(key string) error {
	i.db.QueryRow(`delete from images where id = $1 or name = $1`, key)
	return nil
}

func (i *imageUseCase) DeleteAllUnusedOnes() string {
	i.db.QueryRow(`delete from images where not exists(select image_id from containers c where c.image_id = $1)`)

	return "Total reclaimed space: 0B"
}

func (i *imageUseCase) GetOne(key string) (domain.ImageDetailedResponse, error) {
	var response domain.ImageDetailedResponse
	i.db.QueryRow(`select * from images i where i.id = $1 or i.name = $1`, key).Scan(&response)

	return response, nil
}
