CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) UNIQUE NOT NULL,
    password CHAR(60) NOT NULL,
    CHECK (LENGTH(username) BETWEEN 4 AND 30),
    CHECK (LENGTH(password) = 60)
);

