package main

import (
	twtextparse "github.com/myl7/twitter-text-parse-go"
	"log"
)

func main() {
	res, err := twtextparse.Parse("test")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
