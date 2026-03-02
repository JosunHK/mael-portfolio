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

-- name: UpdateDesktopModeIfNull :exec
UPDATE thumb_mode
SET desktop_id = (SELECT id FROM animation WHERE active = TRUE LIMIT 1)
WHERE desktop_id IS NULL OR (SELECT active FROM animation WHERE id = desktop_id) = FALSE;

-- name: UpdateMobileModeIfNull :exec
UPDATE thumb_mode
SET mobile_id = (SELECT id FROM animation WHERE active = TRUE LIMIT 1)
WHERE mobile_id IS NULL OR (SELECT active FROM animation WHERE id = mobile_id) = FALSE;

-- name: InsertThumbModeIfNotExists :exec
INSERT INTO thumb_mode(
    mobile_id, 
    desktop_id
    ) VALUES (
         (SELECT id FROM animation WHERE active = TRUE LIMIT 1), 
        (SELECT id FROM animation WHERE active = TRUE LIMIT 1));
       
-- name: GetThumbModeRowCount :one
SELECT COUNT(*) FROM thumb_mode;

