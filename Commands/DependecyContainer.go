package Commands

import (
	"github.com/codegangsta/inject"
}

// Dependency Container
type DependencyContainer struct {
	Db    *Model.MongoLoggerRepository `inject`
}

func GetInjector() inject.Injector {
	injector := inject.New()
	injector.Map(&Model.MongoLoggerRepository{"127.0.0.1","versionControl","logs"})
	return injector
}

func GetContainer(container interface{}) error {
	injector := GetInjector()
	return injector.Apply(container)
}

