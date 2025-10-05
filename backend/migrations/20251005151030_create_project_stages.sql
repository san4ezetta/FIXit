-- +goose Up
-- +goose StatementBegin
CREATE TABLE project_stages (
                                id bigserial PRIMARY KEY,
                                name text NOT NULL,
                                created_at timestamptz NOT NULL DEFAULT now(),
                                updated_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO project_stages (name) VALUES
                                      ('Планирование'),
                                      ('Фундамент'),
                                      ('Каркас'),
                                      ('Внешняя отделка'),
                                      ('Внутренняя отделка'),
                                      ('Благоустройство'),
                                      ('Сдача');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM project_stages
WHERE name IN ('Планирование','Фундамент','Каркас','Внешняя отделка','Внутренняя отделка','Благоустройство','Сдача');
DROP TABLE IF EXISTS project_stages;
-- +goose StatementEnd