-- +goose Up
-- +goose StatementBegin
CREATE TABLE defect_comments (
  id bigserial PRIMARY KEY,
  defect_id bigint NOT NULL REFERENCES defects(id),
  author_id bigint NOT NULL REFERENCES users(id),
  body text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_defect_comments_defect_id ON defect_comments(defect_id);
CREATE INDEX idx_defect_comments_author_id ON defect_comments(author_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS defect_comments;
-- +goose StatementEnd
