-- +goose Up
-- +goose StatementBegin
CREATE TABLE blacklist
(
  id serial PRIMARY KEY,
  phone varchar NOT NULL,
  name varchar NOT NULL,
  reason varchar NOT NULL,
  time varchar NOT NULL,
  uploader varchar NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blacklist
-- +goose StatementEnd
