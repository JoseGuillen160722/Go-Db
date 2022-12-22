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
