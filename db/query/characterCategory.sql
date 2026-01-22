-- name: GetCharCat :many
SELECT * FROM character_category
WHERE active = TRUE
ORDER BY sort_order;
