-- +migrate Up
INSERT INTO permission_role (id, role_id, permission_id) VALUES (1, 1, 1);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (2, 1, 2);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (3, 1, 3);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (4, 1, 4);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (5, 2, 5);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (6, 2, 6);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (7, 2, 7);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (8, 2, 8);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (9, 3, 9);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (10, 3, 10);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (11, 3, 11);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (12, 3, 12);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (13, 4, 13);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (14, 4, 14);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (15, 4, 15);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (16, 4, 16);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (17, 5, 17);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (18, 5, 18);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (19, 5, 19);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (20, 5, 20);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (21, 6, 21);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (22, 6, 22);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (23, 6, 23);
INSERT INTO permission_role (id, role_id, permission_id) VALUES (24, 6, 24);

-- +migrate Down
DELETE FROM permission_role WHERE id BETWEEN 1 AND 24;