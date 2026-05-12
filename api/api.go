package api

import (
	"demo/struct/config"
)

func Api() string {
	config := config.NewConfig()
	return config.Key
}
