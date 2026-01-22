-- name: GetCharProfile :many
SELECT * FROM character_profile AS cp
WHERE active = TRUE
AND category_id = ?
ORDER BY sort_order;
