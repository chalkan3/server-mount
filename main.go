package main

import (
	MountControler "./Aplication/Controller/Mount"
	Entity "./Domain/Entites"
	Helpers "./Domain/Helpers"
	Server "./Infrastructure/Server"

	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(Entity.NewConfig)
	container.Provide(Helpers.ParseTemplates)
	container.Provide(MountControler.NewMountController)
	container.Provide(Server.NewServer)

	return container
}

func main() {
	container := buildContainer()
	err := container.Invoke(func(server *Server.ServerConfig) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
