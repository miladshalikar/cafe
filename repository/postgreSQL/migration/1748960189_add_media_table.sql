-- +migrate Up
CREATE TABLE IF NOT EXISTS media
(
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(191) NOT NULL,
    size INTEGER NOT NULL,
    path VARCHAR(191) NOT NULL,
    mime_type VARCHAR(191) NOT NULL,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    bucket VARCHAR(191) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );

-- +migrate Down
DROP TABLE media;