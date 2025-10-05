-- +goose Up
-- +goose StatementBegin
CREATE TABLE defects (
                         id bigserial PRIMARY KEY,
                         project_id bigint NOT NULL REFERENCES projects(id),
                         reporter_id bigint NOT NULL REFERENCES users(id),
                         assignee_id bigint REFERENCES users(id),
                         title text NOT NULL,
                         description text,
                         priority defect_priority NOT NULL DEFAULT 'medium',
                         status defect_status NOT NULL DEFAULT 'new',
                         due_date date,
                         created_at timestamptz NOT NULL DEFAULT now(),
                         updated_at timestamptz NOT NULL DEFAULT now(),
                         closed_at timestamptz,
                         canceled_reason text
);

CREATE INDEX idx_defects_project_id ON defects(project_id);
CREATE INDEX idx_defects_reporter_id ON defects(reporter_id);
CREATE INDEX idx_defects_assignee_id ON defects(assignee_id);
CREATE INDEX idx_defects_status ON defects(status);
CREATE INDEX idx_defects_priority ON defects(priority);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS defects;
-- +goose StatementEnd