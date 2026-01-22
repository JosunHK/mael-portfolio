-- +goose Up
-- +goose StatementBegin
CREATE TABLE character_profile(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    category_id BIGINT NOT NULL,
    label varchar(150) character set utf8 NOT NULL,
    character_desc varchar(150) character set utf8 NOT NULL,
    sort_order INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,

    FOREIGN KEY (category_id) 
    REFERENCES character_category(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE character_profile;
-- +goose StatementEnd
