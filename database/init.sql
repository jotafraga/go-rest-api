USE tcc2020;

CREATE TABLE products 
(
    product_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL,
    description VARCHAR(200) NOT NULL,
    price DOUBLE NOT NULL,
    amount INT NOT NULL,
    PRIMARY KEY ( product_id )
);

INSERT INTO products (name, description, price, amount) VALUES ('Notebook', 'Intel Core i7', 4.999, 17);
INSERT INTO products (name, description, price, amount) VALUES ('SmartPhone', 'Octacore', 1.999, 12);
