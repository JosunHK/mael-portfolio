package cms

import (
	"fmt"
	"mael/cmd/database"
	cmsStruct "mael/cmd/struct/cms"
	resError "mael/cmd/struct/error"
	sqlc "mael/db/generated"
	cmsTemplates "mael/web/templates/contents/cms"
	errorTemplate "mael/web/templates/contents/errorAlert"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type AnimationPatch func(echo.Context) *resError.Error
type AnimationPatchResBody func(echo.Context) (templ.Component, *resError.Error)
type AnimationPatchResFunc func(echo.Context, templ.Component, *resError.Error) error

func GetAnimations(c echo.Context) (templ.Component, *resError.Error) {
	queries := sqlc.New(database.DB)
	res, err := queries.GetAnimations(c.Request().Context())
	resThumb, errThumb := queries.GetThumbMode(c.Request().Context())
	if err != nil {
		return cmsTemplates.AnimationsTable([]sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Failed to retrive Data ", err.Error())
	}
	if errThumb != nil {
		return cmsTemplates.AnimationsTable([]sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Failed to retrive Thumb Data ", errThumb.Error())
	}
	

	return cmsTemplates.AnimationsTable(res, resThumb), nil
}



func DeleteAnimation(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)
	thumbMode, err := queries.GetThumbMode(c.Request().Context())
	if err != nil {
		return resError.New("Failed to retrive Thumb Data ", err.Error())
	}
	err = queries.DeleteAnimation(c.Request().Context(), id)
	if err != nil {
		return resError.New("Failed to Delete Record ", err.Error())
	}

	err = queries.UpdateMobileModeIfNull(c.Request().Context())
	if err != nil {
		return resError.New("Failed to update mobile mode if null ", err.Error())
	}

	err = queries.UpdateDesktopModeIfNull(c.Request().Context())
	if err != nil {
		return resError.New("Failed to update desktop mode if null ", err.Error())
	}
	
	err = DeleteAlert(c, thumbMode, id)
	
	return nil
}

func AddAnimation(c echo.Context) *resError.Error {
	var req = cmsStruct.AddAnimationReq{}

	err := c.Request().ParseForm()
	if err != nil {
		log.Error(fmt.Errorf("failed to parse form %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = database.SQLDecoder().Decode(&req, c.Request().PostForm); err != nil {
		log.Error(fmt.Errorf("Failed to decode / validate Animation Request %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = ValidateAddAnimationReq(req); err != nil {
		log.Error(fmt.Errorf("Failed to validate Aniomation Request %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = insertAnimation(c, req); err != nil {
		log.Error(fmt.Errorf("Failed to insert Record %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = EnsureThumbModeExist(c); err != nil {
		log.Error(fmt.Errorf("Failed to ensure thumb mode exists after creating new animation %v", err))
		return resError.New("Failed to ensure thumb mode exists after creating new animation ", err.Error())
	}

	return nil
}

func insertAnimation(c echo.Context, req cmsStruct.AddAnimationReq) error {
	queries := sqlc.New(database.DB)
	_, err := queries.AddAnimation(c.Request().Context(), req.Label)

	if err != nil {
		return err
	}

	return nil
}

func ValidateAddAnimationReq(req cmsStruct.AddAnimationReq) error {

	if len(req.Label) > 150 {
		return fmt.Errorf("Label is too long")
	}

	return nil
}

func OrderUp(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)
	err = queries.AnimationOrderUp(c.Request().Context(), id)
	if err != nil {
		return resError.New("Failed to Reorder Record ", err.Error())
	}
	return nil
}

func OrderDown(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)
	err = queries.AnimationOrderDown(c.Request().Context(), id)
	if err != nil {
		return resError.New("Failed to Reorder Record ", err.Error())
	}
	return nil
}

func GetAnimtionDetail(c echo.Context) (templ.Component, *resError.Error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return cmsTemplates.AnimationDetail(sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)

	errThumb := EnsureThumbModeExist(c); 
	if errThumb != nil {
		return cmsTemplates.AnimationDetail(sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Failed to ensure thumb mode exists ", errThumb.Error())
	}

	res, err := queries.GetAnimationById(c.Request().Context(), id)
	resThumb, errThumb := queries.GetThumbMode(c.Request().Context())
	if err != nil {
		return cmsTemplates.AnimationDetail(sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Failed to retrive Data ", err.Error())
	}
	if errThumb != nil {
		return cmsTemplates.AnimationDetail(sqlc.Animation{}, sqlc.ThumbMode{}), resError.New("Failed to retrive Thumb Data ", errThumb.Error())
	}
	
	return cmsTemplates.AnimationDetail(res, resThumb), nil
}

func ModifyDetail(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	var req = cmsStruct.ModifyAnimationReq{}

	if err = EnsureThumbModeExist(c); err != nil {
		log.Error(fmt.Errorf("Failed to ensure thumb mode exists after modifying animation %v", err))
		return resError.New("Failed to ensure thumb mode exists after modifying animation ", err.Error())
	}

	err = c.Request().ParseForm()
	if err != nil {
		log.Error(fmt.Errorf("failed to parse form %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = database.SQLDecoder().Decode(&req, c.Request().PostForm); err != nil {
		log.Error(fmt.Errorf("Failed to decode / validate modify Request %v", err))
		return resError.New("Failed to create New Record ", err.Error())
	}

	if err = ValidateAnimationDetail(req); err != nil {
		log.Error(fmt.Errorf("Failed to validate modify request %v", err))
		return resError.New("Failed update animation record ", err.Error())
	}

	res, err := saveAnimation(c, id)
	if err != nil {
		log.Error(fmt.Errorf("Failed save animation frames %v", err))
		return resError.New("Failed save animation frames ", err.Error())
	}

	newAnimation := sqlc.ModifyAnimationParams{
		ID:            id,
		Label:         req.Label,
		AnimationDesc: req.Desc,
		Fps:           req.Fps,
		FramesCount:   res.FramesCount,
		Width:         res.Width,
		Height:        res.Height,
	}

	if err = modifyAnimation(c, newAnimation); err != nil {
		log.Error(fmt.Errorf("Failed update animation record %v", err))
		return resError.New("Failed update animation record ", err.Error())
	}

	return nil
}

func ValidateAnimationDetail(detail cmsStruct.ModifyAnimationReq) error {
	if len(detail.Label) > 150 {
		return fmt.Errorf("Label is too long")
	}

	fps := detail.Fps
	if fps.Valid && (0 > fps.Int32 || fps.Int32 > 1000) {
		return fmt.Errorf("Invalid fps value")
	}

	desc := detail.Desc
	if len(desc) > 1000 {
		return fmt.Errorf("Invalid fps value")
	}

	return nil
}

func modifyAnimation(c echo.Context, params sqlc.ModifyAnimationParams) error {
	queries := sqlc.New(database.DB)
	err := queries.ModifyAnimation(c.Request().Context(), params)
	if err != nil {
		return err
	}

	return nil
}

func ModifyThumbMobile(c echo.Context) *resError.Error{
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	if err = EnsureThumbModeExist(c); err != nil {
		log.Error(fmt.Errorf("Failed to ensure thumb mode exists when modifying mobile mode %v", err))
		return resError.New("Failed to ensure thumb mode exists when modifying mobile mode ", err.Error())
	}

	queries := sqlc.New(database.DB)
	err = queries.UpdateMobileModeIfNull(c.Request().Context())
	if err != nil {
		return resError.New("Failed to update mobile mode if null ", err.Error())
	}

	err = queries.ModifyThumbMobile(c.Request().Context(), sql.NullInt64{Valid: true, Int64: id})
	if err != nil {
		return resError.New("Failed to Reorder Record ", err.Error())
	}
	return nil
}

func ModifyThumbDesktop(c echo.Context) *resError.Error{
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}
	
	if err = EnsureThumbModeExist(c); err != nil {
		log.Error(fmt.Errorf("Failed to ensure thumb mode exists when modifying desktop mode %v", err))
		return resError.New("Failed to ensure thumb mode exists when modifying desktop mode ", err.Error())
	}
	
	queries := sqlc.New(database.DB)
	err = queries.UpdateDesktopModeIfNull(c.Request().Context())
	if err != nil {
		return resError.New("Failed to update desktop mode if null ", err.Error())
	}

	err = queries.ModifyThumbDesktop(c.Request().Context(), sql.NullInt64{Valid: true, Int64: id})
	if err != nil {
		return resError.New("Failed to Reorder Record ", err.Error())
	}
	return nil
}

func EnsureThumbModeExist(c echo.Context) error {
	queries := sqlc.New(database.DB)

	rows, err := queries.GetThumbModeRowCount(c.Request().Context())
	if err != nil {
		return fmt.Errorf("Failed to Get Thumb Mode Row Count: %w", err)
	}
	
	if rows == 0 {
		err := queries.InsertThumbModeIfNotExists(c.Request().Context())
		if err != nil {
			return fmt.Errorf("Failed to Insert Initial Thumb Mode: %w", err)
		}
		return nil
	}
	return nil
}

func DeleteAlert(c echo.Context, thumbModeId sqlc.ThumbMode, id int64) error {
	if thumbModeId.DesktopID.Int64 == id && thumbModeId.MobileID.Int64 == id  {
		deleteThumbAlert := errorTemplate.WarningAlert("Deleted Both Thumbnail Options")
		return deleteThumbAlert.Render(c.Request().Context(), c.Response().Writer)
	}
	if thumbModeId.DesktopID.Int64 == id {
		deleteThumbAlert := errorTemplate.WarningAlert("Deleted the Desktop Thumbnail Option")
		return deleteThumbAlert.Render(c.Request().Context(), c.Response().Writer)
	}
	if thumbModeId.MobileID.Int64 == id {
		deleteThumbAlert := errorTemplate.WarningAlert("Deleted the Mobile Thumbnail Option")
		return deleteThumbAlert.Render(c.Request().Context(), c.Response().Writer)
	}
	return nil
}



