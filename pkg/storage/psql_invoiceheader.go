package storage

import (
	"database/sql"
	"fmt"

	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceheader"
)

const (
	psqlCreateInvoiceHeader = `CREATE TABLE IF NOT EXISTS InvoiceHeaders(id SERIAL NOT NULL,
		ClientName VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT InvoiceHeaders_id_pk PRIMARY KEY(id)
		)`

	psqlInsertInvoiceHeader = `INSERT INTO invoiceheaders(clientname) VALUES ($1) RETURNING id, created_at`
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

// CreateTx implementa la interfaz de invoiceHeader.Storage
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlInsertInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.Id, &m.CreatedAt)
}
