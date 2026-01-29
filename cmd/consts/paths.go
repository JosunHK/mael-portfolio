package consts

import (
	"os"
)

func GetUploadPath() string {
	if os.Getenv("env") == "production" {
		return "railwayAssets"
	} else {
		return "assets"
	}

}
