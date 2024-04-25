-- +goose Up

CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE -- This is a self-referencing foreign key when user gets deleted, all the feeds associated with that user will also be deleted
);

-- +goose Down
DROP TABLE feeds;