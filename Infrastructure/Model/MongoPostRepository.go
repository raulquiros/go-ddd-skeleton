package Model

import (
	"go-ddd-structure/Domain/Types"
)

type MongoPostRepository struct {
	Host       string
	Port       int
	Database   string
	Collection string
}

func (mongoLoggerRepository MongoPostRepository) Create(post Types.Post) (bool, error) {

	sessionDB, err := GetMongoSessionDatabase(mongoLoggerRepository)

	if err != nil {
		return false, err
	}

	if err := sessionDB.C(mongoLoggerRepository.Collection).Insert(post); err != nil {
		return false, err
	}

	return true, nil
}

func (mongoLoggerRepository MongoPostRepository) Update(post Types.Post) (bool, error) {

	//TODO
	return false, nil
}

func (mongoLoggerRepository MongoPostRepository) Delete(post Types.Post) (bool, error) {

	//TODO
	return false, nil
}
