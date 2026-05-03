package api

import (
	"demo/struct/config"
	"fmt"
)

func api() {
	config := config.NewConfig()
	fmt.Println(config.Key)
}
