-- name: CreateTempTable :exec
CREATE TEMPORARY TABLE bar (LIKE foo);

-- name: SelectOne :one
SELECT 1;