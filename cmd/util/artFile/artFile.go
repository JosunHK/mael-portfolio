package artfile

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	artFileStruct "mael/cmd/struct/artFile"
	"os"
	"path/filepath"
)

func leftPad(num int, totalLength int) string {
	
		fileNum := fmt.Sprint("0000", num)
		strNum := fileNum[len(fileNum)-4:]
		return strNum
}

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

func imagePath(fileName string, length int) []string{
	imgList := []string{}
	for i := range length{
	paths := filepath.Join("/static/assets/",fileName,"/",fileName+"_")
		num := leftPad(i+1, 4)
		imgPath := fmt.Sprint(paths, num, ".jpeg")
		imgList = append(imgList, imgPath)
		
	}
		return imgList
}

func GetArtFile(fileName string) artFileStruct.ArtFile {
	artFiles := artFileStruct.ArtFile{
		Name:   fileName,
		Length: imageLength(fileName),
		Height: imageHeight(fileName),
		Path:   imagePath(fileName, imageLength(fileName)),
	}


	return artFiles
}