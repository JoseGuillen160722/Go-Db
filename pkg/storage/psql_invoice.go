package storage

import (
	"database/sql"
	"fmt"

	"github.com/JoseGuillen160722/Go-Db/pkg/invoice"
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceheader"
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceitem"
)

// Usado para trabajar con Postgres - Invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItem   invoiceitem.Storage
}

// PsqlInvoice retorna un nuevo punto de PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItem:   i,
	}
}

// Implementa la interfaz invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	fmt.Printf("Factura creada con el id: %d \n", m.Header.Id)
	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Producto registrados: %d \n", len(m.Items))
	if err := p.storageItem.CreateTx(tx, m.Header.Id, m.Items); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
