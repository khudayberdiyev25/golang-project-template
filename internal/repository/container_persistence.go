package repository

import "database/sql"

type ContainerSQL struct {
	Id      int
	Name    sql.NullString
	ImageId int
	Command sql.NullString
	Created sql.NullTime
	Status  int
}
