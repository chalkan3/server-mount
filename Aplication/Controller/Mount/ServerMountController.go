package Controller

import (
	"html/template"
	"net/http"
	"strconv"

	BLL "../../../Domain/Bussinesses"
)

type MountController struct {
	template *template.Template
}

func (template *MountController) ShowCreatePage(w http.ResponseWriter, r *http.Request) {
	template.template.ExecuteTemplate(w, "main.html", nil)
}

func (template *MountController) ProcessForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		mount := MapFormToStruct(r)
		mount.CreateServer()

	}
}

func MapFormToStruct(r *http.Request) *BLL.Mount {
	var greenBlue bool = false
	var clientPort int = 0
	var apiV2Port int = 0
	var apiAplicacaoPort int = 0

	nome := r.FormValue("nome")
	path := r.FormValue("path")
	conexaoBanco := r.FormValue("BancoConexao")
	bancoUsername := r.FormValue("BancoUsername")
	bancoPassword := r.FormValue("BancoPassword")
	bancoDatabase := r.FormValue("Banco")
	greenBlueForm := r.FormValue("greenblue")
	rabbitVirtualHost := r.FormValue("vhost")
	rabbitUsername := r.FormValue("RabbitUsername")
	rabbitPassword := r.FormValue("RabbitPassword")

	clientPort, _ = strconv.Atoi(r.FormValue("clientport"))
	apiV2Port, _ = strconv.Atoi(r.FormValue("apiv2port"))
	apiAplicacaoPort, _ = strconv.Atoi(r.FormValue("apiv1port"))

	if greenBlueForm == "true" {
		greenBlue = true
	}

	mount := &BLL.Mount{
		Name:              nome,
		Path:              path,
		GreenBlue:         greenBlue,
		ConexaoBanco:      conexaoBanco,
		BancoUsername:     bancoUsername,
		BancoPassword:     bancoPassword,
		BancoDatabase:     bancoDatabase,
		RabbitVirtualHost: rabbitVirtualHost,
		RabbitUsername:    rabbitUsername,
		RabbitPassword:    rabbitPassword,
		ClientPort:        clientPort,
		ApiV2Port:         apiV2Port,
		ApiAplicacaoPort:  apiAplicacaoPort,
	}

	return mount
}

func NewMountController(_template *template.Template) *MountController {
	return &MountController{template: _template}
}
