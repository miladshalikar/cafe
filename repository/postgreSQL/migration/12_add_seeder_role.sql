-- +migrate Up
INSERT INTO roles (id, title) VALUES (1, 'manage-permission');
INSERT INTO roles (id, title) VALUES (2, 'manage-category');
INSERT INTO roles (id, title) VALUES (3, 'manage-item');
INSERT INTO roles (id, title) VALUES (4, 'manage-order');
INSERT INTO roles (id, title) VALUES (5, 'manage-role');
INSERT INTO roles (id, title) VALUES (6, 'manage-user');

-- +migrate Down
DELETE FROM roles WHERE id BETWEEN 1 AND 7;