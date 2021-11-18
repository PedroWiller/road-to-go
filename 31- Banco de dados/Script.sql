CREATE DATABASE DevBook;

USE DevBook;

CREATE TABLE usuarios(
	id int AUTO_INCREMENT PRIMARY KEY,
	nome varchar(50) NOT NULL, 
	email varchar(50) NOT NULL
)

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang'

GRANT ALL PRIVILEGES  ON DevBook.* TO 'golang'@'localhost'