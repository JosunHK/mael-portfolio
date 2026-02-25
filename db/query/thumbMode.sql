-- name: GetThumbMode :one
SELECT * FROM thumb_mode;

-- name: GetThumbMobile :one
SELECT mobile_id FROM thumb_mode;

-- name: GetThumbDesktop :one
SELECT desktop_id FROM thumb_mode;

-- name: ModifyThumbDesktop :exec
UPDATE thumb_mode
SET desktop_id = ?;

-- name: ModifyThumbMobile :exec
UPDATE thumb_mode
SET mobile_id = ?;

