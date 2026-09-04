CREATE TYPE cuisine_type AS ENUM (
    'japanese',
    'chinese',
    'italian',
    'russian',
    'american',
    'european',
    'other'
);

CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    city VARCHAR(255) NOT NULL,
    cuisine cuisine_type NOT NULL,
    callback_url VARCHAR(500) NOT NULL,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);