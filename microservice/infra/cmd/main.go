package main

import (
	"log"
	"microservice/application"
	usecase "microservice/application/use-case"
	"microservice/infra/database"
	webserver "microservice/infra/http"
	"microservice/infra/http/controller"
	"microservice/infra/queues"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseAdapter := database.NewDatabaseAdapter()
	defer databaseAdapter.Close()

	balanceRepository := database.NewBalanceRepository(databaseAdapter.GetDB())
	createBalanceUseCase := usecase.NewCreateBalanceUseCase(balanceRepository)
	findByIdUseCase := usecase.NewFindByIdUseCase(balanceRepository)
	findUseCase := usecase.NewFindUseCase(balanceRepository)

	balanceUseCases := application.NewBalanceUseCases(createBalanceUseCase, findByIdUseCase, findUseCase)

	queueAdapter := queues.NewQueueAdapter()
	balanceHandler := queues.NewBalanceHandler(balanceUseCases)
	consumer := queues.NewConsumer(queueAdapter, []string{"balances"}, balanceHandler)
	consumer.Start()

	controller := controller.NewBalanceController(balanceUseCases)

	webserver := webserver.NewWebServer(":3003")
	webserver.Router.Post("/balance", controller.CreateBalance)
	webserver.Router.Get("/balance/{id}", controller.FindById)
	webserver.Router.Get("/balance", controller.Find)

	log.Println("Server is running on port :3003")
	webserver.Start()
}
