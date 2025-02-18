CREATE TABLE users(
    uid SERIAL NOT NULL,
    email varchar(255) NOT NULL,
    login varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY(uid)
);
INSERT INTO users (email, login, password) VALUES 
('user1@example.com', 'test1', 'test1'),
('user2@example.com', 'test2', 'test2'),
('user3@example.com', 'test3', 'test3');