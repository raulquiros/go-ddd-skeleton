package Controllers

import (
	"github.com/codegangsta/inject"
	"go-ddd-structure/Infrastructure/Model"
)

// Dependency Container
type DependencyContainer struct {
	Db *Model.MongoLoggerRepository `inject`
}

func GetInjector() inject.Injector {
	injector := inject.New()
	injector.Map(&Model.MongoLoggerRepository{"127.0.0.1", "blogtest", "posts"})
	return injector
}

func GetContainer(container interface{}) error {
	injector := GetInjector()
	return injector.Apply(container)
}
