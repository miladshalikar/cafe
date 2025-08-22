-- +migrate Up
CREATE TABLE IF NOT EXISTS payments
(
    id SERIAL PRIMARY KEY,
    order_id INT,
    amount NUMERIC(12,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    method VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_order
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE SET NULL
    );

-- +migrate Down
DROP TABLE IF EXISTS payment;