-- +goose Up
CREATE TABLE visits (
    id int AUTO_INCREMENT NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    primary key (id)
);

-- +goose Down
DROP TABLE visits;
