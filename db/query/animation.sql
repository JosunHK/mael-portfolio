-- name: GetAnimations :many
SELECT * FROM animation
WHERE active = TRUE
ORDER BY sort_order;

-- name: AddAnimation :execresult
INSERT INTO animation(
    label,
    animation_desc
) VALUES (
    ?, ?
);

-- name: DeleteAnimation :exec
UPDATE animation 
SET active = FALSE 
WHERE id = ?;
