package main

import (
	"fmt"

	"github.com/zs5460/jmc"
)

func main() {
	var cfg map[string]string
	jmc.MustLoadConfig("config.json", &cfg)

	fmt.Println(cfg["password"])
	//output: 123456
}
