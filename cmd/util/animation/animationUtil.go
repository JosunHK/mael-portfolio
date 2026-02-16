package animationUtil

import (
	"fmt"
	"mael/cmd/consts"
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
