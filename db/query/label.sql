-- name: GetLabel :one
SELECT label FROM label
WHERE label_key = ?
and active = TRUE;
