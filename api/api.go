package api

import (
	"bytes"
	"demo/struct/config"
	"demo/struct/file"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Api() string {
	config := config.NewConfig()
	return config.Key
}

func Create(fileFlag, nameFlag *string) {
	if *fileFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --file is required with --create")
		os.Exit(1)
	}
	if *nameFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --name is required with --create")
		os.Exit(1)
	}
	data, err := file.ReadFile(*fileFlag)
	if err != nil {
		fmt.Println("Ошибка чтения файла, должен быть json:", err)
		os.Exit(1)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", Api())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ сервера:", string(body))
}

func Update(fileFlag, idFlag *string) {
	if *fileFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --file is required with --update")
		os.Exit(1)
	}

	if *idFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --id is required with --update")
		os.Exit(1)
	}
	data, err := file.ReadFile(*fileFlag)
	if err != nil {
		os.Exit(1)
	}

	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+*idFlag, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", Api())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ сервера:", string(body))

}

func Delete(idFlag *string) {
	if *idFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --id is required with --delete")
		os.Exit(1)
	}
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+*idFlag, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", Api())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ сервера:", string(body))

}

func Get(idFlag *string) {
	if *idFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: --id is required with --get")
		os.Exit(1)
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+*idFlag, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", Api())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ сервера:", string(body))

}
