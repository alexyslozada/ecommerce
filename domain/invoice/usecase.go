package invoice

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/alexyslozada/ecommerce/model"
)

// Invoice implements UseCase
type Invoice struct {
	storage                    Storage
	storageInvoiceDetailReport StorageInvoiceDetailReport
}

// New returns a new Invoice
func New(s Storage, sidr StorageInvoiceDetailReport) Invoice {
	return Invoice{storage: s, storageInvoiceDetailReport: sidr}
}

// Create creates a model.Invoice
func (i Invoice) Create(po *model.PurchaseOrder) error {
	if err := po.Validate(); err != nil {
		return fmt.Errorf("invoice: %w", err)
	}

	invoice, invoiceDetails, err := invoiceFromPurchaseOrder(po)
	if err != nil {
		return fmt.Errorf("%s %w", "invoiceFromPurchaseOrder()", err)
	}

	err = i.storage.Create(&invoice, invoiceDetails)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

func (i Invoice) GetByUserID(userID uuid.UUID) (model.InvoicesReport, error) {
	invoicesHead, err := i.storageInvoiceDetailReport.HeadsByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("invoice: %w", err)
	}

	var invoicesReport model.InvoicesReport
	for _, invoiceHead := range invoicesHead {
		invoiceDetails, err := i.storageInvoiceDetailReport.AllDetailsByInvoiceID(invoiceHead.Invoice.ID)
		if err != nil {
			return nil, fmt.Errorf("%s %w", "storageInvoiceDetail.AllDetailsByInvoiceID()", err)
		}

		invoiceHead.InvoiceDetailsReport = invoiceDetails
		invoicesReport = append(invoicesReport, invoiceHead)
	}

	return invoicesReport, nil
}

// GetAll returns a model.Invoices according to filters and sorts
func (i Invoice) GetAll() (model.InvoicesReport, error) {
	invoices, err := i.storageInvoiceDetailReport.AllHead()
	if err != nil {
		return nil, fmt.Errorf("invoice: %w", err)
	}

	var invoicesReport model.InvoicesReport
	for _, v := range invoices {
		invoiceDetails, err := i.storageInvoiceDetailReport.AllDetailsByInvoiceID(v.Invoice.ID)
		if err != nil {
			return nil, fmt.Errorf("%s %w", "storageInvoiceDetailReport.AllDetailsByInvoiceID()", err)
		}

		v.InvoiceDetailsReport = invoiceDetails
		invoicesReport = append(invoicesReport, v)
	}

	return invoicesReport, nil
}

func invoiceFromPurchaseOrder(po *model.PurchaseOrder) (model.Invoice, model.InvoiceDetails, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	invoice := model.Invoice{
		ID:              ID,
		UserID:          po.UserID,
		PurchaseOrderID: po.ID,
		CreatedAt:       time.Now().Unix(),
	}

	var products model.ProductToPurchases
	err = json.Unmarshal(po.Products, &products)
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	var invoiceDetails model.InvoiceDetails
	for _, v := range products {
		detailID, err := uuid.NewUUID()
		if err != nil {
			return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
		}

		detail := model.InvoiceDetail{
			ID:        detailID,
			InvoiceID: invoice.ID,
			ProductID: v.ProductID,
			Amount:    v.Amount,
			UnitPrice: v.UnitPrice,
			CreatedAt: time.Now().Unix(),
		}

		invoiceDetails = append(invoiceDetails, detail)
	}

	return invoice, invoiceDetails, nil
}
