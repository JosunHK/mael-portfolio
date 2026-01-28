package cmsUtil

import (
	"database/sql"
	"fmt"
	"mael/cmd/database"
	"mael/cmd/struct/cms"
	"mael/cmd/struct/error"
	"mael/db/generated"
	"mael/web/templates/contents/cms"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type AnimationPatch func(echo.Context) *resError.Error
type AnimationPatchResBody func(echo.Context) (templ.Component, *resError.Error)
type AnimationPatchResFunc func(echo.Context, templ.Component, *resError.Error) error

func GetAnimtions(c echo.Context) (templ.Component, *resError.Error) {
	queries := sqlc.New(database.DB)
	res, err := queries.GetAnimations(c.Request().Context())
	if err != nil {
		return cmsTemplates.AnimationsTable([]sqlc.Animation{}), resError.New("Failed to retrive Data ", err.Error())
	}
	return cmsTemplates.AnimationsTable(res), nil
}

func DeleteAnimation(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)
	err = queries.DeleteAnimation(c.Request().Context(), id)
	if err != nil {
		return resError.New("Failed to Delete Record ", err.Error())
	}
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
		return cmsTemplates.AnimationDetail(sqlc.Animation{}), resError.New("Invalid id ", err.Error())
	}

	queries := sqlc.New(database.DB)
	res, err := queries.GetAnimationById(c.Request().Context(), id)
	if err != nil {
		return cmsTemplates.AnimationDetail(sqlc.Animation{}), resError.New("Failed to retrive Data ", err.Error())
	}

	return cmsTemplates.AnimationDetail(res), nil
}

func ModifyDetail(c echo.Context) *resError.Error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return resError.New("Invalid id ", err.Error())
	}

	var req = cmsStruct.ModifyAnimationReq{}

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

	newAnimation := sqlc.ModifyAnimationParams{
		Label:         req.Label,
		AnimationDesc: req.Desc,
		Fps:           req.Fps,
		ID:            id,
	}

	count, err := savesAnimationReturnCount(c, id)
	if err != nil {
		log.Error(fmt.Errorf("Failed save animation frames %v", err))
		return resError.New("Failed save animation frames ", err.Error())
	}

	if count > 0 {
		newAnimation.FramesCount = sql.NullInt32{Valid: true, Int32: int32(count)}
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
