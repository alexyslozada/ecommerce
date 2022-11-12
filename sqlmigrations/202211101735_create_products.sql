CREATE TABLE products (
	id UUID NOT NULL,
	product_name VARCHAR(128) NOT NULL,
	price NUMERIC(10,2) NOT NULL,
	images JSONB NOT NULL,
	description TEXT NOT NULL,
	features JSONB NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT products_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE products IS 'Storage the products for the e-commerce';
