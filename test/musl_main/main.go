package main

import (
	"log"

	twtextparse "github.com/myl7/twitter-text-parse-go/pkg/musl"
)

func main() {
	res, err := twtextparse.Parse("test")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
