-- name: ListSentences :many
SELECT * FROM sentence ORDER BY created_at;

-- name: CreateSentence :one
INSERT INTO sentence (content) VALUES (@content) RETURNING *;

-- name: BatchCreateSentences :copyfrom
INSERT INTO sentence (content) VALUES (@content);

-- name: BatchGetSentencesByIds :many
SELECT * FROM sentence WHERE id = ANY(@id::int[]);

-- name: ListScoredSentences :many
SELECT id, created_at, content, similarity(content, @prompt) AS sml
  FROM sentence
  WHERE content % @prompt
  ORDER BY sml DESC, content;
