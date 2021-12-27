package simple

import (
	"errors"
)

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository // dependency
}

// PROVIDER

// used to create SimpleRepository
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: true,
	}
}

// used to create SimpleService, using parameter dependency
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
