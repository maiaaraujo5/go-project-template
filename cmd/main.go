package main

import (
	"log"

	"github.com/maiaaraujo5/go-project-template/internal/app/project/fx/server"
	"github.com/maiaaraujo5/gostart/application"
)

func main() {
	err := application.Run(
		application.Start.
			WithEcho().
			WithCustomProvider(server.Module()).
			Build())

	if err != nil {
		log.Fatalln(err)
	}
}
