-- +goose Up
-- +goose StatementBegin
CREATE TABLE sub_animation(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    main_id BIGINT NOT NULL,
    label varchar(150) character set utf8 NOT NULL,
    width INT,
    height INT,
    frames_count INT,
    fps INT,
    sort_order INT,
    active BOOLEAN NOT NULL DEFAULT TRUE,

    FOREIGN KEY (main_id) 
    REFERENCES animation(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sub_animation;
-- +goose StatementEnd
