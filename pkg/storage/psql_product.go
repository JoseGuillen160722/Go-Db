package storage

import (
	"database/sql"
	"fmt"

	"github.com/JoseGuillen160722/Go-Db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS Producto(id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observation VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT Producto_id_pk PRIMARY KEY(id)
		)`
	psqlInsertProduct  = `INSERT INTO Producto(name, observation, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	psqlGetallProducts = `SELECT id, name, observation, price, created_at, updated_at FROM Producto`
	psqlGetProductById = `SELECT id, name, observation, price, created_at, updated_at FROM Producto WHERE id = $1`
	psqlUpdateProduct  = `UPDATE Producto SET name = $1, observation = $2, price=$3, updated_at=$4 WHERE id = $5`
	psqlDeleteProduct  = `DELETE FROM Producto WHERE id = $1`
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

// Función que devuelve todos los registros
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetallProducts)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

// Función para búsqueda por Id
func (p *PsqlProduct) GetById(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))

}

func scanRowProduct(s scanner) (*product.Model, error) {

	m := product.Model{}
	ObservationNull := sql.NullString{}
	UpdatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.Id,
		&m.Name,
		&ObservationNull,
		&m.Price,
		&m.CreatedAt,
		&UpdatedAtNull,
	)

	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = ObservationNull.String
	m.UpdatedAt = UpdatedAtNull.Time

	return &m, nil

}

// Fnción para actualizar un registro
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.Id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No existe el Id: %d", m.Id)
	}

	fmt.Println("Se actualizó el producto correctamente")
	return nil

}

func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Se eliminó el producto correctamente")
	return nil
}
