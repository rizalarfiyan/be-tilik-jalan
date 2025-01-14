package main

import (
	"fmt"

	"github.com/rizalarfiyan/be-tilik-jalan/config"
)

func init() {
	config.Init()
}

func main() {
	conf := config.Get()
	fmt.Println(conf)
	fmt.Println("Hello, World!")
}
