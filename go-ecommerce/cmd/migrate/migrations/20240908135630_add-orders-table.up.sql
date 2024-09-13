CREATE TABLE orders (
  id SERIAL,
  total INTEGER NOT NULL,
  status VARCHAR(255) NOT NULL,
  address TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  user_id BigInt NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users (id)
);
