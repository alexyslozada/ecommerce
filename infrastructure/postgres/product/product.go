package product

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alexyslozada/ecommerce/infrastructure/postgres"
	"github.com/alexyslozada/ecommerce/model"

	"github.com/google/uuid"
)

const table = "products"

var fields = []string{
	"id",
	"product_name",
	"price",
	"images",
	"description",
	"features",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

// Product struct that implement the interface domain.product.Storage
type Product struct {
	db *pgxpool.Pool
}

// New returns a new Product storage
func New(db *pgxpool.Pool) Product {
	return Product{db}
}

// Create creates a model.Product
func (p Product) Create(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

// Update this method updates a model.Product by id
func (p Product) Update(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a model.Product by id
func (p Product) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetByID gets an ordered model.Product with filters
func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return p.scanRow(row)
}

// GetAll gets all model.Products with Fields
func (p Product) GetAll() (model.Products, error) {
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Products
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (p Product) scanRow(s pgx.Row) (model.Product, error) {
	m := model.Product{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.ProductName,
		&m.Price,
		&m.Images,
		&m.Description,
		&m.Features,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	return m, nil
}
