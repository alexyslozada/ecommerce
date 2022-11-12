CREATE TABLE purchase_orders (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	products JSONB NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT purchase_orders_id_pk PRIMARY KEY (id),
	CONSTRAINT purchase_orders_user_id_fk FOREIGN KEY (user_id)
	    REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE purchase_orders IS 'Storage the purchase orders';
