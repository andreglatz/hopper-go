-- +goose Up
-- +goose StatementBegin
ALTER TABLE links ADD COLUMN clicks INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE links DROP COLUMN clicks;
-- +goose StatementEnd
