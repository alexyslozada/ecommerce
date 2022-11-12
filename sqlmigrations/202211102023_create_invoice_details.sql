CREATE TABLE invoice_details (
	id UUID NOT NULL,
	invoice_id UUID NOT NULL,
	product_id UUID NOT NULL,
	amount INTEGER NOT NULL,
	unit_price NUMERIC(10,2) NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT invoice_details_id_pk PRIMARY KEY (id),
	CONSTRAINT invoice_details_invoice_id_fk FOREIGN KEY (invoice_id)
            REFERENCES invoices (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT invoice_details_product_id_fk FOREIGN KEY (product_id)
            REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE invoice_details IS 'Storage the details of the invoices';
