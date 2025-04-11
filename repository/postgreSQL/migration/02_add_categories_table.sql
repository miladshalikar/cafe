-- +migrate Up
CREATE TABLE IF NOT EXISTS categories
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    logo VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE categories;