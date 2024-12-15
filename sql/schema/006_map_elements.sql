
-- +goose Up
CREATE TABLE map_elements (
    id TEXT PRIMARY KEY,
    map_id TEXT NOT NULL REFERENCES maps (id) ON DELETE CASCADE,
    element_id TEXT NOT NULL REFERENCES elements (id) ON DELETE CASCADE,
    x INT,
    y INT
);

-- +goose Down
DROP TABLE map_elements;
