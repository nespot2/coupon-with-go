package main

import (
	"coupon-with-go/configs"
	"flag"
	"fmt"
)

func main() {

	var profile string

	flag.StringVar(&profile, "profile", "local", "it is used in the config file")
	flag.Parse()

	cfg, error := configs.New(profile)

	if error != nil {
		panic("죽어!")
	}

	fmt.Println(cfg)
}
