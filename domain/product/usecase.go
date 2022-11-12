package product

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/alexyslozada/ecommerce/model"
)

// Product implements UseCase
type Product struct {
	storage Storage
}

// New returns a new Product
func New(s Storage) Product {
	return Product{storage: s}
}

// Create creates a model.Product
func (p Product) Create(m *model.Product) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID
	if len(m.Images) == 0 {
		m.Images = []byte(`[]`)
	}
	if len(m.Features) == 0 {
		m.Features = []byte(`{}`)
	}

	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a model.Product by id
func (p Product) Update(m *model.Product) error {
	if !m.HasID() {
		return fmt.Errorf("product: %w", model.ErrInvalidID)
	}

	if len(m.Images) == 0 {
		m.Images = []byte(`[]`)
	}
	if len(m.Features) == 0 {
		m.Features = []byte(`{}`)
	}
	m.UpdatedAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a model.Product by id
func (p Product) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	product, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	return product, nil
}

// GetAll returns a model.Products according to filters and sorts
func (p Product) GetAll() (model.Products, error) {
	products, err := p.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("product: %w", err)
	}

	return products, nil
}
