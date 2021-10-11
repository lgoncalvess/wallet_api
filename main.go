package main

import (
	"fmt"
	"wallet/config"
)

func init() {
	config.Load()
}

func main() {
	fmt.Println("running api on Port: ", config.Port)
}
