package Bussinesses

import (
	"strconv"
	"strings"

	Helpers "../../Domain/Helpers"
)

type Mount struct {
	Path          string
	Name          string
	ConexaoBanco  string
	BancoUsername string
	BancoPassword string
	BancoDatabase string

	RabbitVirtualHost string
	RabbitUsername    string
	RabbitPassword    string

	ClientPort       int
	ApiV2Port        int
	ApiAplicacaoPort int

	GreenBlue bool
}

func (mount *Mount) CreateServer() {
	path := "./Clientes/" + mount.Name
	Helpers.CreateDirIfNotExist(mount.Name, mount.GreenBlue)
	CreateServicesFile(path, mount)
	createRemoteCliente(path, mount)

}

func mountServicesStrings(mount *Mount) ([]string, []string) {

	if mount.GreenBlue {
		green := []string{
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Aplicacao | GREEN | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/green/core/Aplicacao\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/green/core/Aplicacao/Cogtive.Core.Api.Aplicacao.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n#KESTREL\nEnvironment=KESTREL_HOSTNAME=http://127.0.0.1:" + strconv.Itoa(mount.ApiAplicacaoPort) + "\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Api V2 | GREEN | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/green/core/Api\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/green/core/Api/Cogtive.Core.Aplicacao.V2.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n#KESTREL\nEnvironment=KESTREL_HOSTNAME=http://127.0.0.1:" + strconv.Itoa(mount.ApiV2Port) + "\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Heimdall | GREEN | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/green/core/Heimdall\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/green/core/Aplicacao/Cogtive.Core.Heimdall.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Jaiminho | GREEN | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/green/core/Jaiminho\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/green/core/Aplicacao/Cogtive.Core.Jaiminho.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n#SMTP\nEnvironment=SMTP_SERVER=mail.cogtive.com.br\nEnvironment=SMTP_PORTA=587\nEnvironment=SMTP_CREDENCIAL=false\nEnvironment=SMTP_USUARIO=\nEnvironment=SMTP_SENHA=\nEnvironment=SMTP_EMAIL=\n\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Velma | GREEN | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/green/core/Velma\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/green/core/Aplicacao/Cogtive.Core.Velma.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
		}

		blue := []string{
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Aplicacao | BLUE | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/blue/core/Aplicacao\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/blue/core/Aplicacao/Cogtive.Core.Api.Aplicacao.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n#KESTREL\nEnvironment=KESTREL_HOSTNAME=http://127.0.0.1:" + strconv.Itoa(mount.ApiAplicacaoPort) + "\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Api V2 | BLUE | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/blue/core/Api\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/blue/core/Api/Cogtive.Core.Aplicacao.V2.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n#KESTREL\nEnvironment=KESTREL_HOSTNAME=http://127.0.0.1:" + strconv.Itoa(mount.ApiV2Port) + "\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Heimdall | BLUE | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/blue/core/Heimdall\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/blue/core/Aplicacao/Cogtive.Core.Heimdall.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Jaiminho | BLUE | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/blue/core/Jaiminho\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/blue/core/Aplicacao/Cogtive.Core.Jaiminho.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n#SMTP\nEnvironment=SMTP_SERVER=mail.cogtive.com.br\nEnvironment=SMTP_PORTA=587\nEnvironment=SMTP_CREDENCIAL=false\nEnvironment=SMTP_USUARIO=\nEnvironment=SMTP_SENHA=\nEnvironment=SMTP_EMAIL=\n\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
			"[Unit]\nDescription=Cogtive " + strings.ToUpper(mount.Name) + " | Velma | BLUE | Core .NET running on CentOS 7\n\n[Service]\nWorkingDirectory=/opt/cogtive/" + mount.Name + "/blue/core/Velma\nExecStart=/usr/bin/dotnet /opt/cogtive/" + mount.Name + "/blue/core/Aplicacao/Cogtive.Core.Velma.dll\nRestart=always\nRestartSec=2  # Restart service after 2 seconds if dotnet service crashes\nSyslogIdentifier=dotnet-example\nUser=root\nEnvironment=ASPNETCORE_ENVIRONMENT=Production\n\n\n\n#POSTGRESQL\nEnvironment=POSTGRES_HOSTNAME=" + mount.ConexaoBanco + "\nEnvironment=POSTGRES_USERNAME=" + mount.BancoUsername + "\nEnvironment=POSTGRES_PASSWORD=" + mount.BancoPassword + "\nEnvironment=POSTGRES_DATABASE=" + mount.BancoDatabase + "\nEnvironment=POSTGRES_PORT=5432\n\n#RABBITMQ\nEnvironment=RABBITMQ_VIRTUALHOST=" + mount.RabbitVirtualHost + "\nEnvironment=RABBITMQ_USERNAME=" + mount.RabbitUsername + "\nEnvironment=RABBITMQ_PASSWORD=" + mount.RabbitPassword + "\nEnvironment=RABBITMQ_HOSTNAME=localhost\nEnvironment=RABBITMQ_POST=5672\nEnvironment=RABBIT_QUEUE_EVENTOS=eventos\nEnvironment=RABBIT_QUEUE_HEIMDALL=heimdall\nEnvironment=RABBIT_QUEUE_VELMA=velma\nEnvironment=RABBIT_QUEUE_EMAIL=email\n\n[Install]\nWantedBy=multi-user.target",
		}

		return green, blue
	}
	return []string{}, []string{}
}

//flag para ver onde esta chamando
func mountServicesFiles(path string, mount *Mount, flag bool) ([]string, []string) {
	if !flag {
		greenFiles := []string{
			Helpers.CreateFile("cgtv-"+mount.Name+"-green-api-aplicacao.service", path+"/green"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-green-api-v2.service", path+"/green"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-green-heimdall.service", path+"/green"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-green-jaiminho.service", path+"/green"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-green-velma.service", path+"/green"),
		}
		blueFiles := []string{
			Helpers.CreateFile("cgtv-"+mount.Name+"-blue-api-aplicacao.service", path+"/blue"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-blue-api-v2.service", path+"/blue"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-blue-heimdall.service", path+"/blue"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-blue-jaiminho.service", path+"/blue"),
			Helpers.CreateFile("cgtv-"+mount.Name+"-blue-velma.service", path+"/blue"),
		}
		return greenFiles, blueFiles

	}
	greenFiles := []string{
		path + "/green/cgtv-" + mount.Name + "-green-api-aplicacao.service",
		path + "/green/cgtv-" + mount.Name + "-green-api-v2.service",
		path + "/green/cgtv-" + mount.Name + "-green-heimdall.service",
		path + "/green/cgtv-" + mount.Name + "-green-jaiminho.service",
		path + "/green/cgtv-" + mount.Name + "-green-velma.service",
	}
	blueFiles := []string{
		path + "/blue/cgtv-" + mount.Name + "-blue-api-aplicacao.service",
		path + "/blue/cgtv-" + mount.Name + "-blue-api-v2.service",
		path + "/blue/cgtv-" + mount.Name + "-blue-heimdall.service",
		path + "/blue/cgtv-" + mount.Name + "-blue-jaiminho.service",
		path + "/blue/cgtv-" + mount.Name + "-blue-velma.service",
	}
	return greenFiles, blueFiles

}

func mountShellScripts(path string, mount *Mount, flag bool) ([]string, []string, []string) {
	var ShellScripts = []string{
		`url_git_node="git@bitbucket.org:cogtive/core.git"

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-* Iniciando Deploy CORE *-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"

if [ -d "core" ]; then
	rm -rf core
fi

mkdir core

if [ -d "core_temp" ]; then
	rm -rf core_temp
fi

mkdir core_temp
cd core_temp

git clone --depth=1 -b $1 $url_git_node .


echo "*-*-*-*-*-*-* Limpando  projeto *-*-*-*-*-*-*" 
dotnet clean 

echo "*-*-*-*-*-*-*   Publicando API   *-*-*-*-*-*-*"
cd Cogtive.Core.Api.V2

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Api

echo ""
echo "*-*-*-*-*-*-* Publicando API APLICACAO *-*-*-*-*-*-*"
cd ../Cogtive.Core.Api.Aplicacao

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Aplicacao

mkdir ../../core/Aplicacao/Util
cp -R Util/Arquivos/ ../../core/Aplicacao/Util/Arquivos
chmod 777 ../../core/Aplicacao/Util/Arquivos/Lote.xsd  ../../core/Aplicacao/Util/Arquivos/Pipa.xsd

echo ""
echo "*-*-*-*-*-*  Publicando  HEIMDALL  *-*-*-*-*-*"
cd ../Cogtive.Core.Heimdall

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Heimdall

echo ""
echo "*-*-*-*-*-*  Publicando  JAIMINHO  *-*-*-*-*-*"
cd ../Cogtive.Core.Jaiminho

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Jaiminho

echo ""
echo "*-*-*-*-*-*-*  Publicando VELMA  *-*-*-*-*-*-*"
cd ../Cogtive.Core.Velma

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Velma

for file in ` + " `ls -1` " + `; do zip "$file".zip $file/*; done

cd ..
cd ..
rm -rf core_temp

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-*-* Deploy finalizado *-*-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
	`,
		`url_git_node="git@bitbucket.org:cogtive/client.git"

url_api_csharp="https://` + mount.Name + `.cogtive.com.br/api/v2"
url_api_aplicacao="https://` + mount.Name + `.cogtive.com.br/api"

port_api="` + strconv.Itoa(mount.ClientPort) + `"

url_api_aplicacao_valida=$(echo $url_api_aplicacao | sed -r 's/\//\\\//g')
url_api_csharp_valida=$(echo $url_api_csharp | sed -r 's/\//\\\//g')

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-* Iniciando  Deploy  Client *-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"

if [ -d "client_temp" ]; then
	rm -rf client_temp
fi

mkdir client_temp
cd client_temp

git clone --depth=1 -b $1 $url_git_node .

cd src/common

sed -i.bak -r -e 's/var[[:blank:]]*APIAplicacao[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var APIAplicacao = \"'$url_api_aplicacao_valida'\"/' \
				-e 's/var[[:blank:]]*APIV2[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var APIV2 = \"'$url_api_csharp_valida'\"/' Constants.js

cd ../../
sed -i.bak -r -e 's/var[[:blank:]]*Porta[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var Porta = '$port_api'/' producao.server.js


echo "*-*-*-*-* Restaurando Dependências  *-*-*-*-*"
yarn install

echo "*-*-*-*-*-*-*-*-* Copilando *-*-*-*-*-*-*-*-*"
yarn deploy

cd js
rm -rf assets

cp -R ../src/assets/ ./

cd  ../../

if [ -d "client" ]; then
	rm -rf client
fi

cp -R client_temp/js ./
mv js client

cp -R client_temp/producao.server.js client/server.js

cd client
npm install express

cd ..
rm -rf client_temp

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-*-* Deploy finalizado *-*-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
	`,
	}

	if !flag {
		blue := []string{
			Helpers.CreateFile("core.sh", path+"/blue"),
			Helpers.CreateFile("client.sh", path+"/blue"),
			Helpers.CreateFile("start.sh", path+"/blue"),
		}
		green := []string{
			Helpers.CreateFile("core.sh", path+"/green"),
			Helpers.CreateFile("client.sh", path+"/green"),
			Helpers.CreateFile("start.sh", path+"/green"),
		}

		return ShellScripts, green, blue
	}

	blue := []string{
		path + "/blue/core.sh",
		path + "/blue/client.sh",
		path + "/blue/start.sh",
	}
	green := []string{
		path + "/green/core.sh",
		path + "/green/client.sh",
		path + "/green/start.sh",
	}

	return ShellScripts, green, blue

}

func CreateServicesFile(path string, mount *Mount) {

	if mount.GreenBlue {
		green, blue := mountServicesStrings(mount)
		greenFiles, bluefiles := mountServicesFiles(path, mount, false)
		shellScript, shellGreenPath, shellBluePath := mountShellScripts(path, mount, false)

		for index, greenShellPath := range shellGreenPath {
			Helpers.WriteFile(shellScript[index], greenShellPath)
		}

		for index, blueShellPath := range shellBluePath {
			Helpers.WriteFile(shellScript[index], blueShellPath)

		}

		for index, serviceGreen := range green {
			Helpers.WriteFile(serviceGreen, greenFiles[index])

		}

		for index, serviceBlue := range blue {
			Helpers.WriteFile(serviceBlue, bluefiles[index])
		}

	}

	return

}

func createRemoteCliente(path string, mount *Mount) {

	greenFiles, blueFiles := mountServicesFiles(path, mount, true)
	_, shellGreenPath, shellBluePath := mountShellScripts(path, mount, true)

	command := "sudo mkdir " + mount.Path + "/" + mount.Name + ";"
	command += "sudo mkdir " + mount.Path + "/" + mount.Name + "/blue;"
	command += "sudo mkdir " + mount.Path + "/" + mount.Name + "/green;"
	command += "sudo chown -R igorguedes:igorguedes *"

	posCommand := "sudo mv  " + mount.Path + mount.Name + "/green/cgtv-* /etc/systemd/system;"
	posCommand += "sudo mv " + mount.Path + mount.Name + "/blue/cgtv-* /etc/systemd/system;"
	posCommand += "sudo mv " + mount.Path + mount.Name + " /opt/cogtive/;"
	posCommand += "sudo chmod 777 /opt/cogtive/" + mount.Name + "/blue/*;"
	posCommand += "sudo chmod 777 /opt/cogtive/" + mount.Name + "/green/*;"

	// posCommand += "sudo systemctl enable cgtv-" + mount.Name + "-* ;"
	//enable

	Helpers.SshAndRunCommand(command)
	for _, localPathGreen := range greenFiles {

		Helpers.ScpGO(localPathGreen, mount.Path+mount.Name+"/green")
	}

	for _, localPathBlue := range blueFiles {
		Helpers.ScpGO(localPathBlue, mount.Path+mount.Name+"/blue")
	}

	for _, shellGreen := range shellGreenPath {
		Helpers.ScpGO(shellGreen, mount.Path+mount.Name+"/green")
	}
	for _, shellBlue := range shellBluePath {
		Helpers.ScpGO(shellBlue, mount.Path+mount.Name+"/blue")
	}

	Helpers.SshAndRunCommand(posCommand)

}
