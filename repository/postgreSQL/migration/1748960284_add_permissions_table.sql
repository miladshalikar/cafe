-- +migrate Up
CREATE TABLE IF NOT EXISTS permissions
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(355),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS permissions;