package animationUtil

import (
	"context"
	"fmt"
	"mael/cmd/consts"
	"mael/cmd/database"
	"os"
	"slices"

	log "github.com/sirupsen/logrus"
	"mael/db/generated"
)

func IsLandscape(animation sqlc.Animation) bool {
	if !animation.Height.Valid || !animation.Width.Valid {
		return false
	}

	if animation.Width.Int32 > animation.Height.Int32 {
		return true
	}
	return false
}

func GetAnimationPaths(id int64) []string {
	entries, err := os.ReadDir(fmt.Sprintf("%v/%d/", consts.LocalAnimationPath, id))
	if err != nil {
		log.Info("No animation frames found!")
		return []string{}
	}

	var paths []string
	for _, entry := range entries {
		paths = append(paths, fmt.Sprintf("%v/%d/%v", consts.RequestAnimationPath, id, entry.Name()))
	}

	slices.Sort(paths)

	return paths
}

func GetSubAnimationPaths(id int64) []string {
	entries, err := os.ReadDir(fmt.Sprintf("%v/%d/", consts.LocalSubAnimationPath, id))
	if err != nil {
		log.Info("No animation frames found!")
		return []string{}
	}

	var paths []string
	for _, entry := range entries {
		paths = append(paths, fmt.Sprintf("%v/%d/%v", consts.RequestSubAnimationPath, id, entry.Name()))
	}

	slices.Sort(paths)

	return paths
}

func GetSubAnimations(c context.Context, id int64) []sqlc.SubAnimation {
	queries := sqlc.New(database.DB)
	res, err := queries.GetSubAnimations(c, id)
	if err != nil {
		return []sqlc.SubAnimation{}
	}

	return res
}
