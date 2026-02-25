-- +goose Up
-- +goose StatementBegin
CREATE TABLE thumb_mode(mobile_id INT NOT NULL, desktop_id INT NOT NULL);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE thumb_mode
-- +goose StatementEnd
