CREATE DATABASE IF NOT EXISTS `wakkason`;

CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_username (username)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE memories (
    memory_id INT PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    text TEXT NOT NULL,
    photo_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE messages (
    message_id INT PRIMARY KEY AUTO_INCREMENT,
    notification_id INT NOT NULL,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    photo_url TEXT,
    scheduled_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
)DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
