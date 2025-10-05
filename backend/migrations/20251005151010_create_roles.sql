-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
                       id smallserial PRIMARY KEY,
                       name text NOT NULL
);

INSERT INTO roles (name) VALUES
                             ('engineer'),
                             ('manager'),
                             ('director');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM roles WHERE name IN ('engineer', 'manager', 'director');
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd