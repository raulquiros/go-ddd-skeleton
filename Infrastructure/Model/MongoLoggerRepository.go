package Model

import (
	"go-ddd-structure/Domain/Types"
	"gopkg.in/mgo.v2"
)

type MongoLoggerRepository struct {
	Host       string
	Database   string
	Collection string
}

func (mongoLoggerRepository MongoLoggerRepository) Create(post Types.Post) (bool, error) {

	session, err := mgo.Dial(mongoLoggerRepository.Host)
	defer session.Close()

	if err != nil {
		return false, err
	}

	if err := session.DB(mongoLoggerRepository.Database).C(mongoLoggerRepository.Collection).Insert(post); err != nil {
		return false, err
	}

	return true, nil
}

func (mongoLoggerRepository MongoLoggerRepository) Update(post Types.Post) (bool, error) {

	//TODO
	return false, nil
}

func (mongoLoggerRepository MongoLoggerRepository) Delete(post Types.Post) (bool, error) {

	//TODO
	return false, nil
}
