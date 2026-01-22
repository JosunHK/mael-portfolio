-- +goose Up
CREATE TABLE character_category(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    label varchar(150) character set utf8 NOT NULL,
    cat_desc varchar(150) character set utf8 NOT NULL,
    cat_img_url nvarchar(150) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

-- +goose Down
DROP TABLE character_category;
