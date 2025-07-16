USE eagle_db;
CREATE TABLE IF NOT EXISTS events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    log TEXT
);
CREATE TABLE IF NOT EXISTS api_key (
    id INT AUTO_INCREMENT PRIMARY KEY,
    auth_key TEXT
);
CREATE TABLE IF NOT EXISTS api_key_readlogs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    auth_key_readlogs TEXT
);
INSERT INTO api_key_readlogs (auth_key_readlogs) 
VALUES ("FMEPJFJw1NyiIVHupksDNhSPCotRKtkLElUfmKYXMG0QO6")