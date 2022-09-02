package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Not success", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response = "Hello World"

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Erro ao recuperar o Hello World", err.Error())
		return
	}

	fmt.Println(response)
}
