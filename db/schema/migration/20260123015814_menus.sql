-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu_item (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    menu_collection_id BIGINT NOT NULL,
    label varchar(150) NOT NULL,
    value varchar(150) NOT NULL,
    sort_order INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE menu_collection (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(150) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menu_item;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE menu_collection;
-- +goose StatementEnd
