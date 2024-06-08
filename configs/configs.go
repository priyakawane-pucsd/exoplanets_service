package configs

import (
	"exoplanetservice/controller"
	"exoplanetservice/repository"
	"exoplanetservice/service"
)

type Configuration struct {
	Name       string
	Controller controller.Config
	Service    service.Config
	Repository repository.Config
}
