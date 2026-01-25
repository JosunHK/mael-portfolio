-- +goose Up
-- +goose StatementBegin
CREATE TABLE animation(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    label varchar(150) character set utf8 NOT NULL,
    animation_desc varchar(150) character set utf8 NOT NULL DEFAULT '',
    width INT,
    height INT,
    frames_count INT,
    fps INT,
    sort_order INT,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE animation;
-- +goose StatementEnd
