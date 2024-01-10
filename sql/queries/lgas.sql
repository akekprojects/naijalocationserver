-- name: AddLga :one
INSERT INTO lgas (name, state_id) VALUES ($1, $2)
RETURNING *;


-- name: GetLgas :many
SELECT * FROM lgas;

-- name: GetStateLgas :many
SELECT * FROM lgas
WHERE state_id = $1;