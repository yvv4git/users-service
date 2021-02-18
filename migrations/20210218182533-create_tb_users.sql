-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    `ID` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(25),
    `age` INTEGER
);

-- +migrate Down
DROP TABLE users;