package menuProvider

import (
	"context"

	log "github.com/sirupsen/logrus"
	"mael/cmd/database"
	i18nUtil "mael/cmd/util/i18n"
	sqlc "mael/db/generated"
)

var MENU_CACHE = make(map[string][]sqlc.MenuItem)

func TranslMenu(ctx context.Context, rawMenu []sqlc.MenuItem) []sqlc.MenuItem {
	menu := []sqlc.MenuItem{}
	for _, item := range rawMenu {
		item.Label = i18nUtil.T(ctx, item.Label)
		menu = append(menu, item)
	}

	return menu

}

func GetMenuPlease(ctx context.Context, key string) []sqlc.MenuItem {
	rawMenu := GetRawMenu(ctx, key)
	rawMenu = append(rawMenu, sqlc.MenuItem{Label: "please_select", Value: ""})

	return TranslMenu(ctx, rawMenu)
}

func GetMenu(ctx context.Context, key string) []sqlc.MenuItem {
	rawMenu := GetRawMenu(ctx, key)

	return TranslMenu(ctx, rawMenu)
}

// func GetLabel(ctx context.Context, key string) string {
// 	DB := database.DB
// 	queries := sqlc.New(DB)
//
// 	result, err := queries.GetLabel(ctx, key)
// 	if err != nil {
// 		log.Error("Error getting label : ", err)
// 		return ""
// 	}
//
// 	return result
// }

func GetRawMenu(ctx context.Context, key string) []sqlc.MenuItem {
	DB := database.DB
	queries := sqlc.New(DB)
	result, ok := MENU_CACHE[key]
	if ok {
		return result
	}

	result, err := queries.GetMenu(ctx, key)
	if err != nil {
		log.Error("Error getting menu: ", err)
		MENU_CACHE[key] = result
		return []sqlc.MenuItem{}
	}

	return result
}

func PurgeCache() error {
	MENU_CACHE = make(map[string][]sqlc.MenuItem)
	return nil
}
