-- +migrate Up
INSERT INTO permissions (id, title, description) VALUES (1, 'create-permission', 'ایجاد دسترسی جدید');
INSERT INTO permissions (id, title, description) VALUES (2, 'index-permission', 'مشاهده لیست دسترسی‌ها');
INSERT INTO permissions (id, title, description) VALUES (3, 'delete-permission', 'حذف دسترسی');
INSERT INTO permissions (id, title, description) VALUES (4, 'edit-permission', 'ویرایش دسترسی');
INSERT INTO permissions (id, title, description) VALUES (5, 'create-category', 'ایجاد دسته‌بندی جدید');
INSERT INTO permissions (id, title, description) VALUES (6, 'index-category', 'مشاهده لیست دسته‌بندی‌ها');
INSERT INTO permissions (id, title, description) VALUES (7, 'delete-category', 'حذف دسته‌بندی');
INSERT INTO permissions (id, title, description) VALUES (8, 'edit-category', 'ویرایش دسته‌بندی');
INSERT INTO permissions (id, title, description) VALUES (9, 'create-item', 'ایجاد آیتم جدید');
INSERT INTO permissions (id, title, description) VALUES (10, 'index-item', 'مشاهده لیست آیتم‌ها');
INSERT INTO permissions (id, title, description) VALUES (11, 'delete-item', 'حذف آیتم');
INSERT INTO permissions (id, title, description) VALUES (12, 'edit-item', 'ویرایش آیتم');
INSERT INTO permissions (id, title, description) VALUES (13, 'create-order', 'ایجاد سفارش جدید');
INSERT INTO permissions (id, title, description) VALUES (14, 'index-order', 'مشاهده لیست سفارش‌ها');
INSERT INTO permissions (id, title, description) VALUES (15, 'delete-order', 'حذف سفارش');
INSERT INTO permissions (id, title, description) VALUES (16, 'edit-order', 'ویرایش سفارش');
INSERT INTO permissions (id, title, description) VALUES (17, 'create-role', 'ایجاد نقش جدید');
INSERT INTO permissions (id, title, description) VALUES (18, 'index-role', 'مشاهده لیست نقش‌ها');
INSERT INTO permissions (id, title, description) VALUES (19, 'delete-role', 'حذف نقش');
INSERT INTO permissions (id, title, description) VALUES (20, 'edit-role', 'ویرایش نقش');
INSERT INTO permissions (id, title, description) VALUES (21, 'create-user', 'ایجاد کاربر جدید');
INSERT INTO permissions (id, title, description) VALUES (22, 'index-user', 'مشاهده لیست کاربران');
INSERT INTO permissions (id, title, description) VALUES (23, 'delete-user', 'حذف کاربر');
INSERT INTO permissions (id, title, description) VALUES (24, 'edit-user', 'ویرایش کاربر');


-- +migrate Down
DELETE FROM permissions WHERE id BETWEEN 1 AND 24;