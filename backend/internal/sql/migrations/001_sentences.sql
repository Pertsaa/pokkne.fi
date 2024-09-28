-- +goose Up
CREATE EXTENSION pg_trgm;

SET pg_trgm.similarity_threshold = 0.1;

CREATE TABLE sentence (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone NOT NULL DEFAULT NOW(),
  content text NOT NULL
);

CREATE INDEX sentence_trgm_idx ON sentence USING GIST (content gist_trgm_ops(siglen=32));

-- +goose Down
DROP TABLE sentence;
