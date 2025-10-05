-- +goose Up
-- +goose StatementBegin
CREATE TABLE defect_history (
  id bigserial PRIMARY KEY,
  defect_id bigint NOT NULL REFERENCES defects(id),
  changed_by bigint REFERENCES users(id),
  changed_at timestamptz NOT NULL DEFAULT now(),
  from_status defect_status,
  to_status defect_status,
  from_assignee bigint REFERENCES users(id),
  to_assignee bigint REFERENCES users(id),
  note text
);

CREATE INDEX idx_defect_history_defect_id ON defect_history(defect_id);
CREATE INDEX idx_defect_history_changed_by ON defect_history(changed_by);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS defect_history;
-- +goose StatementEnd
