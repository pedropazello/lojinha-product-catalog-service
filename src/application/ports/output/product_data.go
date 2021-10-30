package output

import "github.com/google/uuid"

type ProductData struct {
	ID          uuid.UUID
	Name        string
	Description string
}
