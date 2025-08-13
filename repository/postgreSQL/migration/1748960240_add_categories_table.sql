-- +migrate Up
CREATE TABLE IF NOT EXISTS categories
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    media_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_media
        FOREIGN KEY (media_id)
        REFERENCES media(id)
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE categories;