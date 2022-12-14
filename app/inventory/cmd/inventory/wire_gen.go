// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"shop/app/inventory/internal/biz"
	"shop/app/inventory/internal/conf"
	"shop/app/inventory/internal/data"
	"shop/app/inventory/internal/server"
	"shop/app/inventory/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDb(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	inventoryRepo := data.NewInventoryRepo(dataData, logger)
	inventoryUsecase := biz.NewInventoryUsecase(inventoryRepo, logger)
	inventoryService := service.NewInventoryService(inventoryUsecase)
	grpcServer := server.NewGRPCServer(confServer, inventoryService, logger)
	httpServer := server.NewHTTPServer(confServer, inventoryService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
