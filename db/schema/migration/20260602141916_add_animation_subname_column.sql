-- +goose Up
-- +goose StatementBegin
ALTER TABLE animation ADD COLUMN sub_name varchar(150) character set utf8 NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE animation DROP COLUMN sub_name;
-- +goose StatementEnd
