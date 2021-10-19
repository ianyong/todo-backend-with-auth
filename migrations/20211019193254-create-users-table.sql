-- +migrate Up
CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    email           VARCHAR UNIQUE NOT NULL,
    hashed_password VARCHAR NOT NULL,
    role            VARCHAR NOT NULL DEFAULT 'user'
);

-- +migrate Down
DROP TABLE users;
