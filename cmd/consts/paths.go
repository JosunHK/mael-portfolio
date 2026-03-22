package consts

import (
	"os"
)

var LocalAnimationPath = GetUploadPath() + "/uploads/animation"
var LocalSubAnimationPath = GetUploadPath() + "/uploads/animation/sub"
var RequestAnimationPath = "/assets/uploads/animation"
var RequestSubAnimationPath = "/assets/uploads/animation/sub"

func GetUploadPath() string {
	if os.Getenv("env") == "production" {
		return "./railwayAssets"
	} else {
		return "./assets"
	}

}
