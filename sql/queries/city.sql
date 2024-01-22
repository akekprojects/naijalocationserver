-- name: AddCity :one
INSERT INTO city ( name,population, state_id) VALUES ($1,$2, $3) 
RETURNING *;


-- name: GetCities :many
SELECT * FROM city;

-- name: GetStateCities :many
SELECT * FROM city
WHERE state_id = $1;