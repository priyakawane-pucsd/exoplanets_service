package main

import (
	"context"
	"exoplanetservice/configs"
	"exoplanetservice/controller"
	"exoplanetservice/logger"
	"exoplanetservice/repository"
	"exoplanetservice/service"
	"flag"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	env  string
	path string
)

func main() {
	flag.StringVar(&env, "env", "dev", "-env=dev")
	flag.StringVar(&path, "path", ".", "-path=./")
	flag.Parse()
	ctx := context.Background()

	configfile := fmt.Sprintf("%s/%s.yaml", path, env)

	var conf configs.Configuration
	err := configs.ReadConfig(ctx, configfile, &conf)
	if err != nil {
		logger.Panic(ctx, "failed to read configs : %v", err)
	}

	// logging config for debug
	bytes, _ := yaml.Marshal(conf)
	logger.Debug(ctx, "\n%s", bytes)

	// starting application
	repo := repository.NewRepository(ctx, &conf.Repository)
	serviceFactory := service.NewServiceFactory(ctx, &conf.Service, repo)
	cntlr := controller.NewController(ctx, &conf.Controller, serviceFactory)
	cntlr.Listen(ctx)
}
