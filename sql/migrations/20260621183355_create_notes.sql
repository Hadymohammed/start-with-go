-- +goose Up
CREATE TABLE notes (
    id         BIGSERIAL    PRIMARY KEY,
    title      TEXT         NOT NULL,
    content    TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE notes;
