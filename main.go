package main

import (
	"fmt"
	"log"

	"github.com/JoseGuillen160722/Go-Db/pkg/product"
	"github.com/JoseGuillen160722/Go-Db/pkg/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProducto := storage.NewPsqlProduct(storage.PostgresPool())
	serviceProducto := product.NewServiceProducto(storageProducto)

	m := &product.Model{
		Name:         "Llave nosa",
		Price:        70,
		Observations: "Grande y dura",
	}

	if err := serviceProducto.CreateProducto(m); err != nil {
		log.Fatalf("product.Create %v", err)
	}

	fmt.Printf("%+v\n", m)
}
