package Controllers

import (
	"github.com/codegangsta/inject"
	"go-ddd-structure/Infrastructure/Model"
	"go-ddd-structure/Infrastructure/Service"
)

// Dependency Container
type DependencyContainer struct {
	PostRepository *Model.MongoPostRepository `inject`
	GenerateId     *Service.UuidGenerate      `inject`
}

func GetInjector() inject.Injector {
	injector := inject.New()
	//TODO: Get variables by dotenv
	injector.Map(&Model.MongoPostRepository{"127.0.0.1", 27017, "blogtest", "posts"})
	injector.Map(&Service.UuidGenerate{})
	return injector
}

func GetContainer(container interface{}) error {
	injector := GetInjector()
	return injector.Apply(container)
}
