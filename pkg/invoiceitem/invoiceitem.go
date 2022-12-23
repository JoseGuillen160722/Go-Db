package invoiceitem

import "time"

// Model of Invoice Item
type Model struct {
	Id              uint
	InvoiceHeaderId uint
	ProductId       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Slice de model
type Models []*Model

type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
}

// Servicio del invoceitem
type Service struct {
	storage Storage
}

// Retorna un puntero de servicio
func NewServiceInvoiceItem(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar producto
func (s *Service) MigrateInvoiceItem() error {
	return s.storage.Migrate()
}
