package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	fileInfo, err := os.Stat(name)
	if os.IsNotExist(err) {
		fmt.Printf("%s dont exist", name)
		return nil, err
	}
	if err != nil {
		fmt.Printf("%s other problems", name)
		return nil, err
	}
	if fileInfo.IsDir() {
		fmt.Printf("%s is dir", name)
		return nil, err
	}

	if !strings.HasSuffix(fileInfo.Name(), ".json") {
		fmt.Println("Not json file")
		return nil, errors.New("Not json file")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
