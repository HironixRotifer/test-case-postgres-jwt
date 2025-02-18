CREATE TABLE users(
    uid SERIAL NOT NULL,
    email varchar(255) NOT NULL,
    login varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    salt varchar(255) NOT NULL,
    PRIMARY KEY(uid)
);
INSERT INTO users (email, login, password, salt) VALUES 
('user1@example.com', 'test1', 'test1', "super"),
('user2@example.com', 'test2', 'test2', "super"),
('user3@example.com', 'test3', 'test3', "super");