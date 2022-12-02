package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	viaCep := make(chan string)
	apiCep := make(chan string)

	go func() {
		req, err := http.NewRequest(http.MethodGet, "http://viacep.com.br/ws/04870060/json/", nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		viaCep <- fmt.Sprintf("Via Cep: %s\n", string(data))
	}()

	go func() {
		req, err := http.NewRequest(http.MethodGet, "https://cdn.apicep.com/file/apicep/04870-060.json", nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		apiCep <- fmt.Sprintf("Api Cep: %s\n", string(data))
	}()

	var cep string
	select {
	case response := <-viaCep:
		cep = response
	case response := <-apiCep:
		cep = response
	case <-time.After(time.Second):
		log.Fatal("Erro de timeout")
	}
	fmt.Println(cep)
}
