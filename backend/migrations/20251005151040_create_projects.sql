-- +goose Up
-- +goose StatementBegin
CREATE TABLE projects (
  id bigserial PRIMARY KEY,
  code text UNIQUE,
  name text NOT NULL,
  description text,
  start_date date,
  end_date date,
  is_active boolean NOT NULL DEFAULT true,
  project_stage_id bigint NOT NULL REFERENCES project_stages(id),
  created_by bigint REFERENCES users(id),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_projects_project_stage_id ON projects(project_stage_id);
CREATE INDEX idx_projects_created_by ON projects(created_by);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
