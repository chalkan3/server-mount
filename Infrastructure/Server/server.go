package server

import (
	"fmt"
	"net/http"

	MountController "../../Aplication/Controller/Mount"
	Entity "../../Domain/Entites"
)

type ServerConfig struct {
	Config          *Entity.Config
	MountController *MountController.MountController
}

func (server *ServerConfig) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.MountController.ShowCreatePage)
	mux.HandleFunc("/create", server.MountController.ProcessForm)

	return mux
}

func (server *ServerConfig) Run() {
	httpServer := &http.Server{
		Addr:    ":" + server.Config.Port,
		Handler: server.Handler(),
	}

	fmt.Println("Listening... localhost PORT" + httpServer.Addr)
	httpServer.ListenAndServe()
}

func NewServer(_config *Entity.Config, _mount *MountController.MountController) *ServerConfig {
	return &ServerConfig{
		Config:          _config,
		MountController: _mount,
	}
}
