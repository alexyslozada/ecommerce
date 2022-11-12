package model

type InvoiceReport struct {
	Invoice              Invoice               `json:"invoice"`
	User                 User                  `json:"user"`
	InvoiceDetailsReport InvoiceDetailsReports `json:"invoice_details_report"`
}

type InvoicesReport []InvoiceReport

type InvoiceDetailsReport struct {
	InvoiceDetail InvoiceDetail `json:"invoice_detail"`
	Product       Product       `json:"product"`
}

type InvoiceDetailsReports []InvoiceDetailsReport
