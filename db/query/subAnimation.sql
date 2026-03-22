-- name: GetSubAnimations :many
SELECT * FROM sub_animation
WHERE active = TRUE
AND main_id = ?
ORDER BY sort_order;

-- name: GetUploadedSubAnimations :many
SELECT * FROM sub_animation
WHERE active = TRUE
AND main_id = ?
AND frames_count >= 1
ORDER BY sort_order;

-- name: GetSubAnimationById :one
SELECT * FROM sub_animation
WHERE active = TRUE
AND id = ?;

-- name: AddSubAnimation :execresult
INSERT INTO sub_animation(
    label,
    main_id,
    sort_order
) VALUES (
    ?,?, COALESCE(
            (SELECT max_sort_order FROM (
                    SELECT MAX(sort_order) AS max_sort_order FROM animation 
                    WHERE active = TRUE
            ) AS derived_table) + 1
        , 0)
);

-- name: DeleteSubAnimation :exec
UPDATE sub_animation 
SET active = FALSE 
WHERE id = ?;

-- name: ReorderSubAnimation :exec
UPDATE sub_animation AS a
JOIN sub_animation AS t ON t.id = ? AND t.active = TRUE
SET a.sort_order = a.sort_order - 1
WHERE a.active = TRUE
  AND a.sort_order >= t.sort_order;

-- name: SubAnimationOrderUp :exec
UPDATE sub_animation AS a
JOIN sub_animation AS b
  ON b.sort_order = a.sort_order - 1
SET a.sort_order = a.sort_order - 1,
    b.sort_order = b.sort_order + 1
WHERE a.id = ?;

-- name: SubAnimationOrderDown :exec
UPDATE sub_animation AS a
JOIN sub_animation AS b
  ON b.sort_order = a.sort_order + 1
SET a.sort_order = a.sort_order + 1,
    b.sort_order = b.sort_order - 1
WHERE a.id = ?;

-- name: ModifySubAnimation :exec
UPDATE sub_animation 
Set label = ?,
    fps = ?,
    frames_count = COALESCE(?, frames_count),
    height = COALESCE(?, height),
    width = COALESCE(?, width)
WHERE id = ?;
