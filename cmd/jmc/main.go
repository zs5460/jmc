// Copyright 2021 zs. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zs5460/jmc"
	"github.com/zs5460/my"
)

var (
	m string
	i string
	p string
)

func main() {
	flag.StringVar(&m, "m", "encode", "mode:encode or decode")
	flag.StringVar(&i, "i", "config.json", "input file name")
	flag.Parse()

	p = os.Getenv("JMC_K")
	log.Println(p)
	if len(p) == 16 || len(p) == 24 || len(p) == 32 {
		jmc.K = p
	}

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
