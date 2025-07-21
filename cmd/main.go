package main

import (
	"flag"
	"log"
	"readwise-list/pkg/readwise"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	readwise := readwise.NewReadwise()
	tag := flag.String("tag", "", "Tag to filter by")
	archived := flag.Bool("archived", false, "Archived")
	flag.Parse()

	if *tag == "" {
		log.Fatal("Tag is required")
	}

	readwise.GetByTag(*tag, archived)
}
