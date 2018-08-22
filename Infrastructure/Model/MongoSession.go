package Model

import (
	"gopkg.in/mgo.v2"
)

func GetMongoSessionDatabase(mongoLoggerRepository MongoPostRepository) (*mgo.Database, error) {

	session, err := mgo.Dial(mongoLoggerRepository.Host + ":" + string(mongoLoggerRepository.Port))
	defer session.Close()

	if err != nil {
		return nil, err
	}
	return session.DB(mongoLoggerRepository.Database), nil
}
