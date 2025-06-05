-- +goose Up
CREATE TABLE devices(
    id BPCHAR PRIMARY KEY,
    device_type  VARCHAR(20),
    status VARCHAR(20),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE devices;