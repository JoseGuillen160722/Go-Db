package storage

import (
	"database/sql"
	"fmt"

	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceitem"
)

const (
	psqlCreateInvoiceItem = `CREATE TABLE IF NOT EXISTS InvoceItem(id SERIAL NOT NULL,
		InvoiceHeaderId INT NOT NULL,
		ProductId INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT InvoiceItem_id_pk PRIMARY KEY(id)
		)`
	psqlInsertInvoiceItem = `INSERT INTO invoceitem(invoiceheaderid, productid) VALUES ($1, $2) RETURNING id, created_at`
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

func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, HeaderId uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlInsertInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(HeaderId, item.ProductId).Scan(&item.Id, &item.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}
