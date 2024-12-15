
-- +goose Up
CREATE TABLE space_elements (
    id TEXT PRIMARY KEY,
    element_id TEXT NOT NULL REFERENCES elements (id) ON DELETE CASCADE,
    space_id TEXT NOT NULL REFERENCES spaces (id) ON DELETE CASCADE,
    x INT NOT NULL,
    y INT NOT NULL
);

-- +goose Down
DROP TABLE space_elements;
