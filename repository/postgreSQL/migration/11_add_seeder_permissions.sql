-- +migrate Up
INSERT INTO permissions (id, title) VALUES (1, 'create-permission');
INSERT INTO permissions (id, title) VALUES (2, 'index-permission');
INSERT INTO permissions (id, title) VALUES (3, 'delete-permission');
INSERT INTO permissions (id, title) VALUES (4, 'edit-permission');
INSERT INTO permissions (id, title) VALUES (5, 'create-category');
INSERT INTO permissions (id, title) VALUES (6, 'index-category');
INSERT INTO permissions (id, title) VALUES (7, 'delete-category');
INSERT INTO permissions (id, title) VALUES (8, 'edit-category');
INSERT INTO permissions (id, title) VALUES (9, 'create-item');
INSERT INTO permissions (id, title) VALUES (10, 'index-item');
INSERT INTO permissions (id, title) VALUES (11, 'delete-item');
INSERT INTO permissions (id, title) VALUES (12, 'edit-item');
INSERT INTO permissions (id, title) VALUES (13, 'create-order');
INSERT INTO permissions (id, title) VALUES (14, 'index-order');
INSERT INTO permissions (id, title) VALUES (15, 'delete-order');
INSERT INTO permissions (id, title) VALUES (16, 'edit-order');
INSERT INTO permissions (id, title) VALUES (17, 'create-role');
INSERT INTO permissions (id, title) VALUES (18, 'index-role');
INSERT INTO permissions (id, title) VALUES (19, 'delete-role');
INSERT INTO permissions (id, title) VALUES (20, 'edit-role');
INSERT INTO permissions (id, title) VALUES (21, 'create-user');
INSERT INTO permissions (id, title) VALUES (22, 'index-user');
INSERT INTO permissions (id, title) VALUES (23, 'delete-user');
INSERT INTO permissions (id, title) VALUES (24, 'edit-user');

-- +migrate Down
DELETE FROM permissions WHERE id BETWEEN 1 AND 24;