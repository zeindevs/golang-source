CREATE TABLE orders_items (
  id SERIAL,
  quantity INTEGER NOT NULL,
  price INTEGER NOT NULL,
  order_id BigInt NOT NULL,
  product_id BigInt NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (order_id) REFERENCES orders (id),
  FOREIGN KEY (product_id) REFERENCES products (id)
);
