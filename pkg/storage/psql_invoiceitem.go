package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateInvoiceItem = `CREATE TABLE IF NOT EXISTS InvoceItem(id SERIAL NOT NULL,
		InvoiceHeaderId INT NOT NULL,
		ProductId INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT InvoiceItem_id_pk PRIMARY KEY(id)
		)`
)

// Psql usado para trabajar con
type PsqlInvoiceItem struct {
	db *sql.DB
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Implementa una migración para crear una nueva tabla
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	fmt.Println("Migración de Invoice Header realizada correctamente")
	return nil
}
