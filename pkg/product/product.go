package product

import "time"

// Model of Product
type Model struct {
	Id           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Slice de model
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
}

// Servicio del producto
type Service struct {
	storage Storage
}

// Retorna un puntero de servicio
func NewServiceProducto(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar producto
func (s *Service) MigrateProducto() error {
	return s.storage.Migrate()
}

// Es usado para crear un producto
func (s *Service) CreateProducto(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}
