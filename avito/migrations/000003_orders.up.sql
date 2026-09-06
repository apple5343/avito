CREATE TYPE order_status AS ENUM (
    'pending',
    'confirmed',
    'rejected',
    'cooking',
    'ready',
    'delivering',
    'delivered',
    'cancelled'
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    restaurant_id INT NOT NULL REFERENCES restaurants(id),
    status order_status NOT NULL,
    delivery_address TEXT NOT NULL,
    comment TEXT,
    total_price DECIMAL(10, 2) NOT NULL CHECK (total_price >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);