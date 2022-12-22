package main

import (
	"golang.org/x/mod/sumdb/storage"
)

func main() {
	storage.NewPostgresDB()
}
