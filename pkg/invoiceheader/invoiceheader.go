package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of InvoiceHeader
type Model struct {
	Id        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Slice de model
type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
}

// Servicio del invoiceheader
type Service struct {
	storage Storage
}

// Retorna un puntero de servicio
func NewServiceInvoiceHeader(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar producto
func (s *Service) MigrateInvoiceHeader() error {
	return s.storage.Migrate()
}
