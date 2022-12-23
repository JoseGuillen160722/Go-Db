package storage

import (
	"database/sql"
	"fmt"

	"github.com/JoseGuillen160722/Go-Db/pkg/product"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS Producto(id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observation VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT Producto_id_pk PRIMARY KEY(id)
		)`
	psqlInsertProduct = `INSERT INTO Producto(name, observation, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
)

// Psql usado para trabajar con
type PsqlProduct struct {
	db *sql.DB
}

// Función para insertar nuevo registro en la base de datos como producto
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlInsertProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.Id)
	if err != nil {
		return err
	}
	fmt.Println("Se ha registrado correctamente el producto")
	return nil
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Implementa una migración para crear una nueva tabla
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	fmt.Println("Migración de Producto realizada correctamente")
	return nil
}
