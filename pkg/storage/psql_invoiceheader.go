package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateInvoiceHeader = `CREATE TABLE IF NOT EXISTS InvoiceHeaders(id SERIAL NOT NULL,
		ClientName VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT InvoiceHeaders_id_pk PRIMARY KEY(id)
		)`
)

// Psql usado para trabajar con
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Implementa una migración para crear una nueva tabla
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateInvoiceHeader)
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
