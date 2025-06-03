-- +migrate Up
CREATE TABLE IF NOT EXISTS order_item
(
    id SERIAL PRIMARY KEY,
    order_id INT,
    item_id INT,
    price DOUBLE PRECISION NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_item
        FOREIGN KEY (item_id)
        REFERENCES items(id)
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS order_item;