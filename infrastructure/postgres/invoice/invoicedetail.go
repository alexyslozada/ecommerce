package invoice

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/alexyslozada/ecommerce/infrastructure/postgres"
	"github.com/alexyslozada/ecommerce/model"
	"github.com/jackc/pgx/v5"
)

const tableDetails = "invoice_details"

var fieldsDetails = []string{
	"id",
	"invoice_id",
	"product_id",
	"amount",
	"unit_price",
	"created_at",
	"updated_at",
}

var (
	psqlInsertDetails = postgres.BuildSQLInsert(tableDetails, fieldsDetails)
)

func (i Invoice) CreateDetailsBulk(tx pgx.Tx, details model.InvoiceDetails) error {
	batch := pgx.Batch{}
	for _, v := range details {
		batch.Queue(
			psqlInsertDetails,
			v.ID,
			v.InvoiceID,
			v.ProductID,
			v.Amount,
			v.UnitPrice,
			v.CreatedAt,
			postgres.Int64ToNull(v.UpdatedAt),
		).Exec(func(ct pgconn.CommandTag) error {
			return nil
		})
	}

	result := tx.SendBatch(context.Background(), &batch)
	defer func() {
		err := result.Close()
		if err != nil {
			log.Printf("couldn't close result batch: %v", err)
		}
	}()

	_, err := result.Exec()
	if err != nil {
		return err
	}

	return nil
}
