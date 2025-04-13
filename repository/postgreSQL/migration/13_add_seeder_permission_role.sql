-- +migrate Up
INSERT INTO permission_role (id, role_id, permission_id) VALUES (1, 1, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (2, 1, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (3, 1, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (4, 1, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (5, 2, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (6, 2, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (7, 2, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (8, 2, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (9, 3, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (10, 3, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (11, 3, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (12, 3, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (13, 4, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (14, 4, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (15, 4, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (16, 4, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (17, 5, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (18, 5, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (19, 5, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (20, 5, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (21, 6, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (22, 6, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (23, 6, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (24, 6, 4);

-- +migrate Down
DELETE FROM permission_role WHERE id BETWEEN 1 AND 24;