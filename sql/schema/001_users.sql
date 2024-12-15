
-- +goose Up

CREATE TYPE user_role AS ENUM ('Admin', 'User');

CREATE TABLE users (
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    avatar_id TEXT,
    role user_role NOT NULL
);

-- +goose Down

DROP TABLE users;
DROP TYPE user_role;

