// storage.NewSqlServerDB()

	// storageProducto := storage.NewPsqlProduct(storage.PostgresPool())
	// serviceProducto := product.NewService(storageProducto)

	// if err := serviceProducto.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate %v", err)
	// }

	// StorageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.PostgresPool())
	// ServiceInvoiceHeader := invoiceheader.NewServiceInvoiceHeader(StorageInvoiceHeader)

	// if err := ServiceInvoiceHeader.MigrateInvoiceHeader(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }

	// StorageInvoiceItem := storage.NewPsqlInvoiceItem(storage.PostgresPool())
	// ServiceInvoceItem := invoiceitem.NewServiceInvoiceItem(StorageInvoiceItem)
	// if err := ServiceInvoceItem.MigrateInvoiceItem(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }