
-- +migrate Up
CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    score INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
