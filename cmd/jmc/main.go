package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zs5460/jmc"
	"github.com/zs5460/my"
)

var (
	m string
	i string
)

func main() {
	flag.StringVar(&m, "mode", "encode", "encode or decode")
	flag.StringVar(&i, "input", "config.json", "")
	flag.Parse()

	cnt, err := my.ReadText(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch m {
	case "encode":
		cnt = jmc.Encode(cnt)
	case "decode":
		cnt, err = jmc.Decode(cnt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("invalid mode,only support encode and decode.")
		os.Exit(1)
	}
	err = my.WriteText(i, cnt)
	if err != nil {
		fmt.Println(err)
	}
}
