CREATE TABLE users(
    uid SERIAL NOT NULL,
    email varchar(255) NOT NULL,
    ip varchar(45),
    refresh_token varchar(255),
    PRIMARY KEY(uid)
);
INSERT INTO users (email, ip, refresh_token) VALUES 
('user1@example.com', '192.168.1.1', 'token123'),
('user2@example.com', '192.168.1.2', 'token456'),
('user3@example.com', '192.168.1.3', 'token789');