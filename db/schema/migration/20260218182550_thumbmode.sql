-- +goose Up
-- +goose StatementBegin
INSERT INTO thumb_mode (Desktop_id, Mobile_id) 
VALUES (3, 3)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
