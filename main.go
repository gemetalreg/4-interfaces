package main

import (
	"demo/struct/bins"
	"demo/struct/storage"
	"fmt"
	"log"
	"time"
)

func main() {
	// Initialize dependencies
	fs := storage.OSFileSystem{}
	binStorage := storage.NewFileStorage(fs)

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
