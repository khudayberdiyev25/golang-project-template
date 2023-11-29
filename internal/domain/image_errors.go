package domain

type ErrImageNotFound struct {
	Err    string
	Status int
}

func (e *ErrImageNotFound) Error() string {
	return e.Err
}

type ErrImageDoesNotExist struct {
	Err    string
	Status int
}

func (e *ErrImageDoesNotExist) Error() string {
	return e.Err
}

type ErrImageInUse struct {
	Err    string
	Status int
}

func (e *ErrImageInUse) Error() string {
	return e.Err
}
