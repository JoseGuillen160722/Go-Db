package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrorIdNotFound = errors.New("El producto no contiene un Id")
)

// Model of Product
type Model struct {
	Id           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Slice de model
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
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

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s \n",
		m.Id, m.Name, m.Observations, m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format(("2006-01-02")))
}

// GetAll es usado para obtener todos los registros de la tabla
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetById es usado para obtener un registro en específico
func (s *Service) GetById(id uint) (*Model, error) {
	return s.storage.GetById(id)
}

func (s *Service) Update(m *Model) error {
	if m.Id == 0 {
		return ErrorIdNotFound
	}
	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
