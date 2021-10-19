-- +migrate Up
CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    email           VARCHAR UNIQUE NOT NULL,
    hashed_password VARCHAR NOT NULL
);

-- +migrate Down
DROP TABLE users;
