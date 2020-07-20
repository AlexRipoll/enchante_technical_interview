DROP DATABASE IF EXISTS enchainte_db;

CREATE DATABASE IF NOT EXISTS enchainte_db;

USE enchainte_db;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_on VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS products (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(13,2) NOT NULL,
    seller_id CHAR(36) NOT NULL,
    created_on VARCHAR(255) NOT NULL,
    updated_on VARCHAR(255) NOT NULL DEFAULT "",
    FOREIGN KEY (seller_id) REFERENCES users (id)
);

DROP TABLE IF EXISTS orders;

CREATE TABLE IF NOT EXISTS orders (
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    product_id CHAR(36) NOT NULL,
    seller_id CHAR(36) NOT NULL,
    price DECIMAL(13,2) NOT NULL,
    quantity TINYINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ,
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (seller_id) REFERENCES users (id)
);

INSERT INTO
    users (id, username, email, password, role, created_on)
    VALUES
     (
        "2217014a-8520-4aec-bed7-dc9b70dffbcf",
        "johnDoe",
        "john.Doe@gmail.com",
        "$2y$10$UZemHdSweFA8Y9h26Ftc6eGVaCJ0gBoKHjO41Vy1oTuIDAsRNubvi",
        "user",
        ""
    ),
     (
        "7d695ec2-8979-4c90-8758-9f57badf5cf4",
        "seller",
        "seller@gmail.com",
        "$2y$10$UZemHdSweFA8Y9h26Ftc6eGVaCJ0gBoKHjO41Vy1oTuIDAsRNubvi",
        "seller",
        ""
    )
;

INSERT INTO
    products (id, name, price, seller_id, created_on, updated_on)
    VALUES
    (
    "44b06cc0-6f2c-47d6-ae46-be5e04153e9c",
     "PS5",
     599.99,
     "7d695ec2-8979-4c90-8758-9f57badf5cf4",
     "",
     ""
     ),
    (
    "dfc7c5da-6b9a-4eb6-864a-0339a3c0d22d",
     "PC",
     1702.00,
     "7d695ec2-8979-4c90-8758-9f57badf5cf4",
     "",
     ""
     )
;
