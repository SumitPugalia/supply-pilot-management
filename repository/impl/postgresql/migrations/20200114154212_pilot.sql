-- +goose Up
-- +goose StatementBegin
CREATE TABLE pilots (id VARCHAR PRIMARY KEY,
	user_id VARCHAR,
	supplier_id VARCHAR,
	market_id VARCHAR,
	service_id VARCHAR,
	code_name VARCHAR,
	state VARCHAR,
	created_at timestamp,
	updated_at timestamp,
	deleted boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pilots;
-- +goose StatementEnd
