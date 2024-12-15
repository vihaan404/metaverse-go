
-- +goose Up
CREATE TABLE maps (
    id TEXT PRIMARY KEY,
    width INT NOT NULL,
    height INT NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE maps;
