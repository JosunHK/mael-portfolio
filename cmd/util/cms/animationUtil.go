package cmsUtil

import (
	"fmt"
	"mael/cmd/consts"
	"os"
	"slices"

	log "github.com/sirupsen/logrus"
)

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
