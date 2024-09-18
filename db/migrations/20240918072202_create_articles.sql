-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id int AUTO_INCREMENT,
    title varchar(100),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles;
-- +goose StatementEnd
