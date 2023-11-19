package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron"
	"github.com/sherpalden/sal-dist/appConfig"
	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/repository"
	"github.com/sherpalden/sal-dist/service"
	"github.com/sherpalden/sal-dist/store"
)

func main() {
	c := cron.New()
	/* sec min hr dayOfMonth month dayOfWeek
	cron specification (0 0 0 1 * *) in Golang schedules a job to run at midnight on the first day of every month.
	*/
	c.AddFunc("0 */1 * * * *", func() {
		env := appConfig.GetEnv()
		client, err := ethclient.DialContext(context.Background(), env.NetworkURl)
		if err != nil {
			log.Fatal("failed to initialize network client")
		}

		sender, err := entity.NewSender(env.OwnerPrivateKey, env.OwnerPublicAddress)
		if err != nil {
			log.Fatal("failed to initialize sender")
		}

		empStore := store.NewJSONFileStore("./employees.json")
		empRepo := repository.NewEmployeeRepository(empStore)
		salaryService := service.NewSalaryService(sender, empRepo)
		if err := salaryService.Distribute(client); err != nil {
			log.Fatalf("failed to distribute salary: %v", err)
		}
	})
	c.Start()
	select {}
}
