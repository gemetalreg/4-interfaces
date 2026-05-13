package main

import (
	"demo/struct/api"
	"demo/struct/bins"
	"demo/struct/storage"
	"flag"
	"fmt"
	"log"
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

	// Dispatch based on command
	switch {
	case *createFlag:
		api.Create(fileFlag, nameFlag)
	case *updateFlag:
		api.Update(fileFlag, idFlag)
	case *deleteFlag:
		api.Delete(idFlag)
	case *getFlag:
		api.Get(idFlag)
	case *listFlag:
		loadedList, err := binStorage.LoadBinList()
		if err != nil {
			log.Fatalf("Failed to load bin list: %v", err)
		}

		fmt.Println("Loaded bins:")
		for _, bin := range loadedList.Bins {
			fmt.Printf("%s: %s\n", bin.Name, bin.Id)
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
