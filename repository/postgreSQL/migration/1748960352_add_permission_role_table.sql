-- +migrate Up
CREATE TABLE IF NOT EXISTS permission_role
(
    id SERIAL PRIMARY KEY,
    role_id INT,
    permission_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_permission
        FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS permission_role;