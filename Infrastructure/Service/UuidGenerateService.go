package Service

import "github.com/satori/go.uuid"

type UuidGenerate struct {
}

func (uuidGenerate UuidGenerate) Generate() (string, error) {

	uuid, err := uuid.NewV4()
	return uuid.String(), err

}
