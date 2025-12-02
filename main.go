package main

import (
	"fmt"

	"github.com/alonsoF100/weather-service/internal/config"
)

func main() {
	config := config.Load()

	fmt.Println(config)
}
