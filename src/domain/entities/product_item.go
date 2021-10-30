package entities

import "github.com/google/uuid"

type ProductItem struct {
	ID         uuid.UUID
	Color      string
	PriceCents int
}
