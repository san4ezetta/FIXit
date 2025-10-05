-- +goose Up
-- +goose StatementBegin
CREATE TYPE defect_status AS ENUM ('new','in_progress','in_review','closed');
CREATE TYPE defect_priority AS ENUM ('low','medium','high','critical');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS defect_priority;
DROP TYPE IF EXISTS defect_status;
-- +goose StatementEnd