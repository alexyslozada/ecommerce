package invoice

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alexyslozada/ecommerce/infrastructure/postgres"
	"github.com/alexyslozada/ecommerce/model"
)

const (
	table = "invoices"
)

var fields = []string{
	"id",
	"user_id",
	"purchase_order_id",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
)

// Invoice struct that implement the interface domain.invoice.Storage
type Invoice struct {
	db *pgxpool.Pool
}

// New returns a new Invoice storage
func New(db *pgxpool.Pool) Invoice {
	return Invoice{db: db}
}

func (i Invoice) getTx() (pgx.Tx, error) {
	return i.db.Begin(context.Background())
}

// Create creates a model.Invoice
func (i Invoice) Create(m *model.Invoice, ms model.InvoiceDetails) error {
	tx, err := i.getTx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.UserID,
		m.PurchaseOrderID,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		errRollback := tx.Rollback(context.Background())
		if errRollback != nil {
			return fmt.Errorf("%s %w", errRollback, err)
		}

		return err
	}

	err = i.CreateDetailsBulk(tx, ms)
	if err != nil {
		errRollback := tx.Rollback(context.Background())
		if errRollback != nil {
			return fmt.Errorf("%s %w", errRollback, err)
		}

		return err
	}

	errCommit := tx.Commit(context.Background())
	if errCommit != nil {
		return errCommit
	}

	return nil
}
