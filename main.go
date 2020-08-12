package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	USERNAME = "API USERNAME"
	PASSWORD = "API PASSWORD"
	URL      = "API URL"
)

func main() {
	// Authentication
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para API . Erro: ", err.Error())
		return
	}

	req.SetBasicAuth(USERNAME, PASSWORD)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina da API. Erro: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da API. Erro: ", err.Error())
			return
		}
		fmt.Println("")
		fmt.Println(string(body))
	}

}
