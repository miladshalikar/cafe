-- +migrate Up
INSERT INTO roles (id, title, description) VALUES (1, 'manage-permission', 'مدیریت دسترسی‌ها');
INSERT INTO roles (id, title, description) VALUES (2, 'manage-category', 'مدیریت دسته‌بندی‌ها');
INSERT INTO roles (id, title, description) VALUES (3, 'manage-item', 'مدیریت آیتم‌ها');
INSERT INTO roles (id, title, description) VALUES (4, 'manage-order', 'مدیریت سفارش‌ها');
INSERT INTO roles (id, title, description) VALUES (5, 'manage-role', 'مدیریت نقش‌ها');
INSERT INTO roles (id, title, description) VALUES (6, 'manage-user', 'مدیریت کاربران');

-- +migrate Down
DELETE FROM roles WHERE id BETWEEN 1 AND 7;