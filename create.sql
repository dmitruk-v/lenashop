-------------------------------------------------------
-- CUSTOMERS
-------------------------------------------------------
DROP TABLE IF EXISTS customer CASCADE;
CREATE TABLE customer (
	customer_id SERIAL PRIMARY KEY,
	email VARCHAR(45) NOT NULL,
	phone VARCHAR(45),
	address VARCHAR(100) NOT NULL,
	token VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

-------------------------------------------------------
-- CATEGORIES
-------------------------------------------------------
DROP TABLE IF EXISTS category CASCADE;
CREATE TABLE category (
	category_id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);

INSERT INTO category (title, description, created_at) VALUES ('Jewelry', 'All sorts of jewelry', NOW());

-------------------------------------------------------
-- PRODUCTS
-------------------------------------------------------
DROP TABLE IF EXISTS product CASCADE;
CREATE TABLE product (
	product_id SERIAL PRIMARY KEY,
	category_id INT REFERENCES category(category_id),
	title VARCHAR NOT NULL,
	price NUMERIC(10,2) NOT NULL,
	quantity INT NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4209-19 Золотистая (4820004209195)', 160, 99, 'Брошь-булавка "Кот", полностью декорирована стразами.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4209-51 Розово-зеленая (4820004209515)', 99, 33, 'Брошь в виде ящерицы, декорирована кристаллами и стразами.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4208-40 Серебристая (4820004208402)', 75, 33, 'Аксессуары ТМ "TRAUM" (Украина) Брошь - стилизованный цветок со стразами, серебристый.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4208-12 Кремовая', 99, 33, '', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4209-50 Разноцветная (4820004209508)', 120, 33, '', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4209-18 Белая (4820004209188)', 160, 19, 'Брошь в виде цветка с перламутровыми лепестками, декорирована стразами и перламутровыми бусинами.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Декоративные булавки Traum 4209-812 2шт Темная медь (7117110000)', 85, 33, 'Комплект больших декоративных булавок в ретро-стиле, украшенных фигуркой феи и стилизованным цветком.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь-булавка Traum 4209-86 Золотистая', 85, 33, '', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь Traum 4208-32 Розовая (4820004208327)', 99, 33, 'Аксессуары ТМ "TRAUM" (Украина) Объемная брошь в форме цветка розового цвета с золотым стеблем.', NOW());
INSERT INTO product (category_id, title, price, quantity, description, created_at) VALUES (1, 'Брошь женская Traum 4209-60 Красная (4820004209607)', 89, 33, 'Женская брошь Traum 4209-60 в виде мака с золотым напылением, покрыта красной и черной эмалью.', NOW());

DROP TABLE IF EXISTS product_image CASCADE;
CREATE TABLE product_image (
	image_id SERIAL PRIMARY KEY,
	product_id INT REFERENCES product(product_id),
	image_url VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

INSERT INTO product_image (product_id, image_url, created_at) VALUES (1, '/assets/images/69964935.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (1, '/assets/images/69964957.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (2, '/assets/images/16482366.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (2, '/assets/images/16482387.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (3, '/assets/images/181113521.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (3, '/assets/images/181113507.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (4, '/assets/images/177648749.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (4, '/assets/images/177648750.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (5, '/assets/images/21923543.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (5, '/assets/images/21923567.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (6, '/assets/images/69965201.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (6, '/assets/images/69965221.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (7, '/assets/images/33562077.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (7, '/assets/images/33562057.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (8, '/assets/images/177648780.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (8, '/assets/images/177648779.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (9, '/assets/images/181113519.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (9, '/assets/images/181113505.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (10, '/assets/images/14641217.jpg', NOW());
INSERT INTO product_image (product_id, image_url, created_at) VALUES (10, '/assets/images/14641240.jpg', NOW());
-------------------------------------------------------
-- ORDERS
-------------------------------------------------------
DROP TABLE IF EXISTS customer_order CASCADE;
CREATE TABLE customer_order (
	order_id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL REFERENCES customer(customer_id),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

DROP TABLE IF EXISTS customer_order_has_product CASCADE;
CREATE TABLE customer_order_has_product (
	order_id INT NOT NULL REFERENCES customer_order(order_id),
  product_id INT NOT NULL REFERENCES product(product_id),
  quantity SMALLINT NOT NULL,
  CONSTRAINT customer_order_has_product_pk PRIMARY KEY(order_id, product_id)
);
-------------------------------------------------------
-- ORDERS
-------------------------------------------------------
DROP TABLE IF EXISTS cart CASCADE;
CREATE TABLE cart (
	cart_id SERIAL NOT NULL UNIQUE,
  customer_id INT NOT NULL REFERENCES customer(customer_id),
	created_at TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS cart_has_product CASCADE;
CREATE TABLE cart_has_product (
	cart_id INT NOT NULL REFERENCES cart(cart_id),
  product_id INT NOT NULL REFERENCES product(product_id),
	quantity INT NOT NULL,
	CONSTRAINT cart_has_product_pk PRIMARY KEY(cart_id, product_id)
);