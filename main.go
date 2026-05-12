package main

import (
	"demo/struct/api"
	"demo/struct/bins"
	"demo/struct/storage"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	createFlag := flag.Bool("create", false, "Create a new bin from file")
	updateFlag := flag.Bool("update", false, "Update a bin by ID")
	deleteFlag := flag.Bool("delete", false, "Delete a bin by ID")
	getFlag := flag.Bool("get", false, "Get a bin by ID")
	listFlag := flag.Bool("list", false, "List all bins")
	fileFlag := flag.String("file", "", "Path to JSON file (required for create/update)")
	idFlag := flag.String("id", "", "Bin ID (required for update/delete/get)")
	nameFlag := flag.String("name", "", "Name for new bin (required for create)")

	// ✅ CRITICAL: Parse the command-line flags
	flag.Parse()

	// Validate: exactly one command must be used
	commands := []*bool{createFlag, updateFlag, deleteFlag, getFlag, listFlag}
	activeCount := 0
	for _, cmd := range commands {
		if *cmd {
			activeCount++
		}
	}

	if activeCount != 1 {
		fmt.Fprintln(os.Stderr, "Error: Exactly one command (--create, --update, --delete, --get, --list) must be specified.")
		flag.Usage()
		os.Exit(1)
	}

	fs := storage.OSFileSystem{}
	binStorage := storage.NewFileStorage(fs)
	key := api.Api()

	// Dispatch based on command
	switch {
	case *createFlag:
		if *fileFlag == "" {
			fmt.Fprintln(os.Stderr, "Error: --file is required with --create")
			os.Exit(1)
		}
		if *nameFlag == "" {
			fmt.Fprintln(os.Stderr, "Error: --name is required with --create")
			os.Exit(1)
		}
		client := &http.Client{}

		req, err := http.NewRequest("UPDATE", "https://api.jsonbin.io/v3/b/", nil)
		if err != nil {
			fmt.Println("Ошибка создания запроса:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Master-Key", key)

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

	case *updateFlag:

		if *idFlag == "" {
			fmt.Fprintln(os.Stderr, "Error: --id is required with --get")
			os.Exit(1)
		}
		client := &http.Client{}

		req, err := http.NewRequest("UPDATE", "https://api.jsonbin.io/v3/b/"+*idFlag, nil)
		if err != nil {
			fmt.Println("Ошибка создания запроса:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Master-Key", key)

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

	case *deleteFlag:
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
		req.Header.Set("X-Master-Key", key)

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

	case *getFlag:
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
		req.Header.Set("X-Master-Key", key)

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

	case *listFlag:
		loadedList, err := binStorage.LoadBinList()
		if err != nil {
			log.Fatalf("Failed to load bin list: %v", err)
		}

		fmt.Println("Loaded bins:")
		for _, bin := range loadedList.Bins {
			fmt.Printf("%s: %s)\n", bin.Name, bin.Id)
		}

	default:
		fmt.Fprintln(os.Stderr, "Unexpected command state.")
		os.Exit(1)
	}

	// Initialize dependencies

	// Create some bins
	bin1 := bins.NewBin("1", false, time.Now(), "First Bin")
	bin2 := bins.NewBin("2", true, time.Now(), "Private Bin")

	// Create and save bin list
	binList := bins.NewBinList()
	binList.Bins = append(binList.Bins, *bin1, *bin2)

	if err := binStorage.SaveBinList(binList); err != nil {
		log.Fatalf("Failed to save bin list: %v", err)
	}

	// Load and display bin list
	loadedList, err := binStorage.LoadBinList()
	if err != nil {
		log.Fatalf("Failed to load bin list: %v", err)
	}

	fmt.Println("Loaded bins:")
	for _, bin := range loadedList.Bins {
		fmt.Printf("- %s (private: %t)\n", bin.Name, bin.Private)
	}
}
