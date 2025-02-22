// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"my-clean-architecture/api"
	"my-clean-architecture/api/handlers"
	"my-clean-architecture/app"
	"my-clean-architecture/repo"
	"my-clean-architecture/service"
)

// Injectors from wire.go:

func InitServer() *app.Server {
	engine := app.NewGinEngine()
	db := app.InitDB()
	iArticleRepo := repo.NewMysqlArticleRepository(db)
	iArticleService := service.NewArticleService(iArticleRepo)
	articleHandler := handlers.NewArticleHandler(iArticleService)
	router := api.NewRouter(articleHandler)
	server := app.NewServer(engine, router)
	return server
}
