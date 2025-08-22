-- 建立 book 表格
CREATE TABLE IF NOT EXISTS book (
    isbn INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 確保使用者存在並有正確權限
CREATE USER IF NOT EXISTS 'bookuser'@'%' IDENTIFIED BY 'bookpass';
GRANT ALL PRIVILEGES ON bookdb.* TO 'bookuser'@'%';
FLUSH PRIVILEGES;

-- 插入測試資料
INSERT INTO book (isbn, name, publisher) VALUES
(0, 'The C Programming Language', 'Prentice Hall'),
(1, 'Design Patterns', 'Addison-Wesley'),
(2, 'Introduction to Algorithms', 'MIT Press'),
(3, 'Clean Code', 'Prentice Hall'),
(4, 'Refactoring', 'Addison-Wesley')
ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    publisher = VALUES(publisher);
