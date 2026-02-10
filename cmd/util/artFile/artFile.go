package artfile

import (
	"image"
	_ "image/jpeg"
	"log"
	artFileStruct "mael/cmd/struct/artFile"
	"os"
	"path/filepath"
)

func imageLength(fileName string) int{

		fileP := filepath.Join("web/static/assets/",fileName,"/")
		file, err := os.ReadDir(fileP)
		if err != nil {
			log.Fatalf("Failed to read directory: %v", err)
		}
		fileLen := len(file)
		return fileLen

}

func imageHeight(fileName string) int{
		fileP := filepath.Join("web/static/assets/",fileName,"/",fileName+"_0001.jpeg")
		file, err := os.Open(fileP)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()
		img, _, err := image.DecodeConfig(file)
		if err != nil {
			log.Fatalf("Failed to decode image: %v", err)
		}
		return img.Height/16
}

func GetArtFile(fileName string) artFileStruct.ArtFile {
	artFiles := artFileStruct.ArtFile{
		Name:   fileName,
		Length: imageLength(fileName),
		Height: imageHeight(fileName),
		Path:   filepath.Join("web/static/assets/", fileName, "/"),
	}


	return artFiles
}