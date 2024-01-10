-- name: GetAllStates :many
SELECT * FROM states;


-- name: GetStatesA :many
SELECT * FROM states
WHERE name LIKE $1 || '%';

-- name: GetState :one
SELECT * FROM states
WHERE name = $1 ;