-- +goose Up
-- +goose StatementBegin
CREATE TABLE defect_attachments (
  id bigserial PRIMARY KEY,
  defect_id bigint NOT NULL REFERENCES defects(id),
  uploaded_by bigint REFERENCES users(id),
  file_name text NOT NULL,
  mime_type text,
  size_bytes bigint,
  storage_path text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_defect_attachments_defect_id ON defect_attachments(defect_id);
CREATE INDEX idx_defect_attachments_uploaded_by ON defect_attachments(uploaded_by);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS defect_attachments;
-- +goose StatementEnd
