package cmsUtil

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"mael/cmd/database"
	"mael/cmd/struct/cms"
	"mael/cmd/struct/error"
	"mael/db/generated"
	"mael/web/templates/contents/cms"
	"strconv"
)

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
		return resError.New("Failed to Delete Record ", err.Error())
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

	if err = schema.NewDecoder().Decode(&req, c.Request().PostForm); err != nil {
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
	_, err := queries.AddAnimation(c.Request().Context(), sqlc.AddAnimationParams{
		Label:         req.Label,
		AnimationDesc: req.Desc,
	})

	if err != nil {
		return err
	}

	return nil
}

func ValidateAddAnimationReq(req cmsStruct.AddAnimationReq) error {

	if len(req.Label) > 150 {
		return fmt.Errorf("Label is too long")
	}

	if len(req.Desc) > 150 {
		return fmt.Errorf("Description is too long")
	}

	return nil
}
