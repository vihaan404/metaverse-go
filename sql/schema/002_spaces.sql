
-- +goose Up
CREATE TABLE spaces (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    width INT NOT NULL,
    height INT,
    thumbnail TEXT,
    creator_id TEXT NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE spaces;
