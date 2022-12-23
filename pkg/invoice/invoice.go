package invoice

import (
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceheader"
	"github.com/JoseGuillen160722/Go-Db/pkg/invoiceitem"
)

// Modelo de factura
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Interfaz que va a implementar el almacenamiento de base de datos
type Storage interface {
	Create(*Model) error
}

// Servicio de una factura
type Service struct {
	storage Storage
}

// Función constructora que regresa un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Crea una nueva factura
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
