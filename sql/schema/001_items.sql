-- +goose Up
CREATE TABLE items(
    id UUID PRIMARY KEY,
    item_type  VARCHAR(20),
    barcode TEXT UNIQUE NOT NULL,
    state VARCHAR(20) NOT NULL,
    location TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE items;