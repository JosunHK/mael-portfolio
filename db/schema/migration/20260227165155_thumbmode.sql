-- +goose Up
-- +goose StatementBegin
CREATE TABLE thumb_mode(
    mobile_id BIGINT NULL, desktop_id BIGINT NULL,
    FOREIGN KEY (mobile_id) REFERENCES animation(id),
    FOREIGN KEY (desktop_id) REFERENCES animation(id)

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS thumb_mode;
-- +goose StatementEnd