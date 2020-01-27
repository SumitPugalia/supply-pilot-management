-- +goose Up
-- +goose StatementBegin
CREATE TABLE pilots (
  id uuid PRIMARY KEY,
  user_id uuid,
  supplier_id uuid,
  market_id uuid,
  service_id uuid,
  code_name TEXT,
  status TEXT,
  created_at timestamp,
  updated_at timestamp,
  deleted boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pilots;
-- +goose StatementEnd
