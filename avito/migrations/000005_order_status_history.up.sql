CREATE TABLE IF NOT EXISTS order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    status order_status NOT NULL,
    changed_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_order_status_history() RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO order_status_history(order_id, status)
    VALUES (NEW.id, NEW.status);
    RETURN NULL;
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS update_order_status_history_trigger ON orders;
CREATE TRIGGER update_order_status_history_trigger AFTER INSERT OR UPDATE ON orders
FOR EACH ROW EXECUTE PROCEDURE update_order_status_history();