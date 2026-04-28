package storage

import (
	"demo/struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveLocalBin(bin *bins.Bin) {

	data, err := json.Marshal(bin)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.WriteFile("bin.json", data, 0644)

}

func SaveLocalBinList(binList *bins.BinList) {

	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.WriteFile("binList.json", data, 0644)

}

func ReadFile() (*bins.BinList, error) {
	var binList bins.BinList
	data, err := os.ReadFile("binList.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	return &binList, nil
}
