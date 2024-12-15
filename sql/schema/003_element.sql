
-- +goose Up
CREATE TABLE elements (
    id TEXT PRIMARY KEY,
    width INT NOT NULL,
    height INT NOT NULL,
    image_url TEXT NOT NULL
);

-- +goose Down
DROP TABLE elements;
