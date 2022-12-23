package main

import (
	"log"

	"github.com/JoseGuillen160722/Go-Db/pkg/invoice"
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceheader"
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceitem"
	"github.com/JoseGuillen160722/Go-Db/pkg/storage"
)

func main() {
	storage.NewPostgresDB()

	// storageProducto := storage.NewPsqlProduct(storage.PostgresPool())
	// serviceProducto := product.NewServiceProducto(storageProducto)

	// ************************* INSERTAR NUEVO REGISTRO EN LA BASE DE DATOS
	// m:= &product.Model{
	// 	Name: "Patata hot",
	// 	Price: 69,
	// 	Observations: "Very hot",
	// }

	// if err:= serviceProducto.CreateProducto(m); err!= nil{
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// ************************ MÉTODO DE BÚSQUEDA POR ID
	// m, err := serviceProducto.GetById(3)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("No hay producto con este id")
	// case err != nil:
	// 	log.Fatalf("product.GetById: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// *********************** MÉTODO DE ACTUALIZACIÓN DE DATOS DE LA BASE DE DATOS

	// m := &product.Model{
	// 	Id:           1,
	// 	Name:         "Brida de dos patas",
	// 	Observations: "Quiero esa brida",
	// 	Price:        13,
	// }
	// err := serviceProducto.Update(m)
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// ****************************** MÉTODO DE ELIMINACIÓN DE UN REGISTRO DE LA BASE DE DATOS
	// err := serviceProducto.Delete(3)
	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	storageHeader := storage.NewPsqlInvoiceHeader(storage.PostgresPool())
	storageItems := storage.NewPsqlInvoiceItem(storage.PostgresPool())
	storageInvoice := storage.NewPsqlInvoice(storage.PostgresPool(), storageHeader, storageItems)
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Doña Mari",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductId: 1},
			&invoiceitem.Model{ProductId: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create %v", err)
	}
}
