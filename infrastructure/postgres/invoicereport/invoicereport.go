package invoicereport

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"

	"github.com/alexyslozada/ecommerce/model"
)

const (
	psqlHeadInvoiceReport = `
	SELECT invoices.id, invoices.purchase_order_id, invoices.created_at,
	       users.id, users.email, users.details, users.created_at
	FROM invoices INNER JOIN users
	    ON invoices.user_id = users.id
	`
	psqlDetailInvoiceReport = `
	SELECT invoice_details.id, invoice_details.amount,
	       invoice_details.unit_price, products.id,
		   products.product_name, products.images,
		   products.description, products.features,
		   products.created_at
	FROM invoice_details INNER JOIN products
	    ON invoice_details.product_id = products.id
	WHERE invoice_details.invoice_id = $1
	`
)

type InvoiceReport struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) InvoiceReport {
	return InvoiceReport{db: db}
}

func (ir InvoiceReport) HeadByInvoiceID(ID uuid.UUID) (model.InvoiceReport, error) {
	row := ir.db.QueryRow(
		context.Background(),
		psqlHeadInvoiceReport+" WHERE invoices.ID = $1",
		ID,
	)

	return ir.scanHead(row)
}

func (ir InvoiceReport) HeadsByUserID(userID uuid.UUID) (model.InvoicesReport, error) {
	rows, err := ir.db.Query(
		context.Background(),
		psqlHeadInvoiceReport+" WHERE users.ID = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp model.InvoicesReport
	for rows.Next() {
		r, err := ir.scanHead(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, r)
	}

	return resp, nil
}

func (ir InvoiceReport) AllHead() (model.InvoicesReport, error) {
	rows, err := ir.db.Query(
		context.Background(),
		psqlHeadInvoiceReport,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp model.InvoicesReport
	for rows.Next() {
		r, err := ir.scanHead(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, r)
	}

	return resp, nil
}

func (ir InvoiceReport) AllDetailsByInvoiceID(ID uuid.UUID) (model.InvoiceDetailsReports, error) {
	rows, err := ir.db.Query(
		context.Background(),
		psqlDetailInvoiceReport,
		ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp model.InvoiceDetailsReports
	for rows.Next() {
		r, err := ir.scanDetail(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, r)
	}

	return resp, nil
}

func (ir InvoiceReport) scanHead(s pgx.Row) (model.InvoiceReport, error) {
	invoice := model.Invoice{}
	user := model.User{}

	err := s.Scan(
		&invoice.ID,
		&invoice.PurchaseOrderID,
		&invoice.CreatedAt,
		&user.ID,
		&user.Email,
		&user.Details,
		&user.CreatedAt,
	)
	if err != nil {
		return model.InvoiceReport{}, err
	}

	r := model.InvoiceReport{
		Invoice: invoice,
		User:    user,
	}

	return r, nil
}

func (ir InvoiceReport) scanDetail(s pgx.Row) (model.InvoiceDetailsReport, error) {
	invoiceDetail := model.InvoiceDetail{}
	product := model.Product{}

	err := s.Scan(
		&invoiceDetail.ID,
		&invoiceDetail.Amount,
		&invoiceDetail.UnitPrice,
		&product.ID,
		&product.ProductName,
		&product.Images,
		&product.Description,
		&product.Features,
		&product.CreatedAt,
	)
	if err != nil {
		return model.InvoiceDetailsReport{}, err
	}

	r := model.InvoiceDetailsReport{
		InvoiceDetail: invoiceDetail,
		Product:       product,
	}

	return r, nil
}
