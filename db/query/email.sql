-- name: GetEmail :one
SELECT * FROM emails
WHERE id = $1 LIMIT 1;

-- name: Listemails :many
SELECT * FROM emails
ORDER BY subjectEmail;

-- name: CreateEmail :one
INSERT INTO emails (
  subjectEmail, messageEmail
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEmail :exec
DELETE FROM emails
WHERE id = $1;

-- name: UpdateEmail :one
UPDATE emails
  set subjectEmail = $2,
  messageEmail = $3
WHERE id = $1
RETURNING *;