package i18n

import (
	"fmt"
	"slices"

	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	i18nStruct "mael/cmd/struct/i18n"
	i18nUtil "mael/cmd/util/i18n"
	responseUtil "mael/cmd/util/response"
	i18nTemplates "mael/web/templates/contents/i18n"
)

var decoder = schema.NewDecoder()

func GetItems(c echo.Context) error {
	locale := c.Param("locale")
	items := i18nUtil.GetItems(locale)
	slices.Reverse(items)
	return responseUtil.HTML(c, i18nTemplates.I18nTableItems(items))
}

func AddItems(c echo.Context) error {
	locale := c.Param("locale")
	err := c.Request().ParseForm()
	if err != nil {
		log.Error(err)
		return err
	}

	var item i18nStruct.Item

	err = decoder.Decode(&item, c.Request().PostForm)
	if err != nil {
		log.Error(err)
		return err
	}

	if !isItemValid(item, i18nUtil.GetItems(locale)) {
		err := fmt.Errorf("Invalid item")
		log.Error(err)
		return err
	}

	if err := i18nUtil.AddItem(locale, item); err != nil {
		log.Error(err)
		return err
	}

	return GetItems(c)
}

// I'll add this when I want
func DeleteItems(c echo.Context) error {
	locale := c.Param("locale")
	err := c.Request().ParseForm()
	if err != nil {
		log.Error(err)
		return err
	}

	var item i18nStruct.Item

	err = decoder.Decode(&item, c.Request().PostForm)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := i18nUtil.DeleteItem(locale, item.Key); err != nil {
		log.Error(err)
		return err
	}

	return GetItems(c)

}

func isItemValid(item i18nStruct.Item, items []i18nStruct.Item) bool {
	if item.Key == "" || item.Default == "" {
		return false
	}

	for _, item := range items {
		if item.One != "" && item.Many == "" {
			return false
		}
	}

	return true
}
