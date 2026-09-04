CREATE TABLE IF NOT EXISTS order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    status order_status NOT NULL,
    changed_at TIMESTAMP NOT NULL DEFAULT NOW()
);