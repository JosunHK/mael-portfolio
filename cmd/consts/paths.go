package consts

import (
	"os"
)

var LocalAnimationPath = GetUploadPath() + "/uploads/animation"
var RequestAnimationPath = "/assets/uploads/animation"

func GetUploadPath() string {
	if os.Getenv("env") == "production" {
		return "./railwayAssets"
	} else {
		return "./assets"
	}

}
