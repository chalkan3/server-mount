package Helpers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/sfreiberg/simplessh"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./Presentation/Templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}

func CreateDirIfNotExist(dir string, blueGreen bool) {
	if _, err := os.Stat("./Clientes/" + dir); os.IsNotExist(err) {
		err = os.MkdirAll("./Clientes/"+dir, 0755)
		if err != nil {
			panic(err)
		}
		if blueGreen {
			if _, err := os.Stat("./Clientes/" + dir + "/blue"); os.IsNotExist(err) {
				err = os.MkdirAll("./Clientes/"+dir+"/blue", 0755)
				if err != nil {
					panic(err)
				}
			}

			if _, err := os.Stat("./Clientes/" + dir + "/green"); os.IsNotExist(err) {
				err = os.MkdirAll("./Clientes/"+dir+"/green", 0755)
				if err != nil {
					panic(err)
				}
			}
		}

	}
}

func WriteFile(payload string, path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(payload)
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}

}

func CreateFile(filename string, dir string) string {
	path := dir + "/" + filename

	var _, err = os.Stat(dir + path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return ""
		}
		defer file.Close()
	}
	fmt.Println("File was created " + filename + " on PATH" + dir)

	return path
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func SshAndRunCommand(command string) error {
	var client *simplessh.Client
	var err error
	var pathKey string = "/Users/igorguedes/Desktop/gcloud3"

	if client, err = simplessh.ConnectWithKeyFile("104.197.1.230", "igorguedes", pathKey); err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.Exec(command); err != nil {
		log.Println(err)
	}

	return nil
}

func getKeyFile() (key ssh.Signer, err error) {
	file := "/Users/igorguedes/Desktop/gcloud3"
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}

func ScpGO(path string, remotePath string) {
	key, err := getKeyFile()
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: "igorguedes",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", "104.197.1.230:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	err = scp.CopyPath(path, remotePath, session)
	if err != nil {
		panic("Failed to Copy: " + err.Error())
	}
	defer session.Close()
}
