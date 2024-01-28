-- +goose Up
-- +goose StatementBegin
CREATE TABLE links(
  id SERIAL PRIMARY KEY,
  short TEXT NOT NULL,
  original TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd
