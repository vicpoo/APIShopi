CREATE DATABASE IF NOT EXISTS Shop;
USE Shop;

CREATE TABLE IF NOT EXISTS users (
    id_usuario INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(50),
    lastname VARCHAR(50),
);

CREATE TABLE IF NOT EXISTS clothes (
    id_clothes INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    size VARCHAR(10),
    price DECIMAL(10,2),
    stock INT DEFAULT 0,
    imagen_url VARCHAR(255)
);