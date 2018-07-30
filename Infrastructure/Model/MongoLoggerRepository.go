package Model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"versioncontrol/Domain/Types"
)

type MongoLoggerRepository struct {
	Host string
	Database string
	Collection string
}

func (mongoLoggerRepository MongoLoggerRepository) Save(logPushRequest Types.LogPushRequest) (bool, error) {

	session, err := mgo.Dial(mongoLoggerRepository.Host)
	defer session.Close()

	if err != nil {
		return false, err
	}

	if err := session.DB(mongoLoggerRepository.Database).C(mongoLoggerRepository.Collection).Insert(logPushRequest); err != nil{
		return false, err
	}

	return true, nil
}


func (mongoLoggerRepository MongoLoggerRepository) GetLog(conditions bson.M) ([]Types.LogPushRequest, error) {

	session, err := mgo.Dial(mongoLoggerRepository.Host)
	defer session.Close()

	if err != nil {
		return nil, err
	}
	var logs []Types.LogPushRequest

	session.DB(mongoLoggerRepository.Database).C(mongoLoggerRepository.Collection).Find(conditions).Sort("-$natural").All(&logs)

	return logs, nil
}