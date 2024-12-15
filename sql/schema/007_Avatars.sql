
-- +goose Up
CREATE TABLE avatars (
    id TEXT PRIMARY KEY,
    image_url TEXT,
    name TEXT
);

-- +goose Down
DROP TABLE avatars;
