-- name: GetAnimations :many
SELECT * FROM animation
WHERE active = TRUE
ORDER BY sort_order;

-- name: GetAnimationById :one
SELECT * FROM animation
WHERE active = TRUE
AND id = ?;

-- name: AddAnimation :execresult
INSERT INTO animation(
    label,
    sort_order
) VALUES (
    ?, COALESCE(
            (SELECT max_sort_order FROM (
                    SELECT MAX(sort_order) AS max_sort_order FROM animation 
                    WHERE active = TRUE
            ) AS derived_table) + 1
        , 0)
);

-- name: DeleteAnimation :exec
UPDATE animation 
SET active = FALSE 
WHERE id = ?;

-- name: ReorderAnimation :exec
UPDATE animation AS a
JOIN animation AS t ON t.id = ? AND t.active = TRUE
SET a.sort_order = a.sort_order - 1
WHERE a.active = TRUE
  AND a.sort_order >= t.sort_order;

-- name: AnimationOrderUp :exec
UPDATE animation AS a
JOIN animation AS b
  ON b.sort_order = a.sort_order - 1
SET a.sort_order = a.sort_order - 1,
    b.sort_order = b.sort_order + 1
WHERE a.id = ?;

-- name: AnimationOrderDown :exec
UPDATE animation AS a
JOIN animation AS b
  ON b.sort_order = a.sort_order + 1
SET a.sort_order = a.sort_order + 1,
    b.sort_order = b.sort_order - 1
WHERE a.id = ?;

-- name: ModifyAnimation :exec
UPDATE animation 
Set label = ?,
    fps = ?,
    animation_desc = ?,
    frames_count = COALESCE(?, frames_count)
WHERE id = ?;
