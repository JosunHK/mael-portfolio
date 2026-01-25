package cmsUtil

import (
	"archive/zip"
	"fmt"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

var AllowedImageExt = []string{
	".jpg",
	".JPG",
	".jpeg",
	".JPEG",
	".png",
	".PNG",
	".webp",
	".WEBP",
}

type decodeFunc func(io.Reader) (image.Image, error)

var imageDecoder = map[string]decodeFunc{
	".jpg":  jpeg.Decode,
	".JPG":  jpeg.Decode,
	".jpeg": jpeg.Decode,
	".JPEG": jpeg.Decode,
	".png":  png.Decode,
	".PNG":  png.Decode,
}

var destPrefixAnimation = "./assets/uploads/animation/"
var destPrefixImages = "./assets/uploads/images/"

func savesAnimation(c echo.Context, id int64) error {
	srcFile, err := c.FormFile("file")
	if err != nil {
		return nil //no file provided
	}

	src, err := srcFile.Open()
	if err != nil {
		return fmt.Errorf("Unable to open file")
	}
	defer src.Close()

	ext := filepath.Ext(srcFile.Filename)
	if ext != ".zip" {
		return fmt.Errorf("Not a .zip file")
	}

	reader, err := zip.NewReader(src, srcFile.Size)
	if err != nil {
		return fmt.Errorf("Failed to create Reader %v", err)
	}

	files := sortZippedFiles(reader)
	if !(len(files) > 0) {
		return fmt.Errorf("Zip contains no images")
	}

	if err = clearAndCreateDir(fmt.Sprintf("%v%d/", destPrefixAnimation, id)); err != nil {
		return fmt.Errorf("Failed to create Reader %v", err)
	}

	index := 0
	for _, file := range files {
		fileExt := filepath.Ext(file.Name)
		if fileExt == "" || !slices.Contains(AllowedImageExt, fileExt) {
			log.Errorf("Invalid file Ext %v ", fileExt)
			continue
		}

		imgFile, err := file.Open()
		if err != nil {
			log.Errorf("Failed to open file %v ", err)
			continue
		}
		defer imgFile.Close()

		err = convertImageAndSaveWebp(index, id, imgFile, fileExt)
		if err != nil {
			log.Errorf("Failed to save file %v ", err)
			continue
		} else {
			index++
		}
	}

	return nil
}

func scaleAndSaveWebp(index int, id int64, file io.ReadCloser) error {
	output, err := os.Create(fmt.Sprintf("%v%d/%v.webp", destPrefixAnimation, id, index))
	if err != nil {
		return fmt.Errorf("Failed to create encoder for image %v ", err)
	}
	defer output.Close()

	img, err := webp.Decode(file, &decoder.Options{})
	if err != nil {
		return fmt.Errorf("Cannot decode webp %v", err)
	}

	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	// write new image to file
	jpeg.Encode(output, m, nil)

	return nil
}

func convertImageAndSaveWebp(index int, id int64, file io.ReadCloser, fileExt string) error {
	if fileExt == ".webp" || fileExt == ".WEBP" {
		return scaleAndSaveWebp(index, id, file)
	}

	decoder := imageDecoder[fileExt]
	if decoder == nil {
		return fmt.Errorf("Cannot find decoder for %v ", fileExt)
	}

	img, err := decoder(file)
	if err != nil {
		return fmt.Errorf("Failed to create decoder %v ", err)
	}

	var scaledImg image.Image
	if img.Bounds().Dx() > img.Bounds().Dy() {
		scaledImg = resize.Resize(1080, 0, img, resize.Lanczos3)
	} else {
		scaledImg = resize.Resize(0, 1080, img, resize.Lanczos3)
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 100)
	if err != nil {
		return fmt.Errorf("Failed add options to lossy encoder %v ", err)
	}

	output, err := os.Create(fmt.Sprintf("%v%d/%v.webp", destPrefixAnimation, id, index))
	if err != nil {
		return fmt.Errorf("Failed to create path for image %v ", err)
	}
	defer output.Close()

	if err := webp.Encode(output, scaledImg, options); err != nil {
		return fmt.Errorf("Failed write image to path %v ", err)
	}

	return nil
}

func sortZippedFiles(reader *zip.Reader) []*zip.File {
	var files []*zip.File
	for _, f := range reader.File {
		files = append(files, f)
	}

	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
	})

	return files
}

func clearAndCreateDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Errorf("Failed to clear folder %v ", err)
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	return nil
}
