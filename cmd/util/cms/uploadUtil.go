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
	"mael/cmd/consts"
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

type DecodeFunc func(io.Reader) (image.Image, error)
type ImageHandler func(string, io.ReadCloser) error

var imageHandlerMap = map[string]ImageHandler{
	".jpg":  convertJpegAndSaveWebp,
	".JPG":  convertJpegAndSaveWebp,
	".jpeg": convertJpegAndSaveWebp,
	".JPEG": convertJpegAndSaveWebp,
	".png":  convertPngAndSaveWebp,
	".PNG":  convertPngAndSaveWebp,
	".webp": scaleAndSaveWebp,
	".WEBP": scaleAndSaveWebp,
}

var destPrefixAnimation = consts.GetUploadPath() + "/uploads/animation/"
var destPrefixImages = consts.GetUploadPath() + "/uploads/images/"

func savesAnimationReturnCount(c echo.Context, id int64) (int, error) {
	srcFile, err := c.FormFile("file")
	if err != nil {
		return -1, nil
	}

	src, err := srcFile.Open()
	if err != nil {
		return 0, fmt.Errorf("Unable to open file")
	}
	defer src.Close()

	ext := filepath.Ext(srcFile.Filename)
	if ext != ".zip" {
		return 0, fmt.Errorf("Not a .zip file")
	}

	reader, err := zip.NewReader(src, srcFile.Size)
	if err != nil {
		return 0, fmt.Errorf("Failed to create Reader %v", err)
	}

	files := sortZippedFiles(reader)
	if !(len(files) > 0) {
		return 0, fmt.Errorf("Zip contains no images")
	}

	if err = clearAndCreateDir(fmt.Sprintf("%v%d/", destPrefixAnimation, id)); err != nil {
		return 0, fmt.Errorf("Failed to create Reader %v", err)
	}

	pathPrefix := fmt.Sprintf("%v%d/", destPrefixAnimation, id)

	return saveUnzippedFilesReturnCount(pathPrefix, files), nil
}

func saveUnzippedFilesReturnCount(path string, files []*zip.File) int {
	index := 0
	for _, file := range files {
		path := fmt.Sprintf("%v%v.webp", path, index)
		if err := saveUnzippedFile(path, file); err != nil {
			log.Errorf("Failed to save unzipped file %v", err)
			continue
		}

		index++
	}
	return index
}

func saveUnzippedFile(path string, file *zip.File) error {
	fileExt := filepath.Ext(file.Name)
	if fileExt == "" || !slices.Contains(AllowedImageExt, fileExt) {
		return fmt.Errorf("Invalid file Ext %v ", fileExt)
	}

	imgFile, err := file.Open()
	if err != nil {
		return fmt.Errorf("Failed to open file %v ", err)
	}
	defer imgFile.Close()

	err = saveImage(path, imgFile, fileExt)
	if err != nil {
		return fmt.Errorf("Failed to save file %v ", err)
	}
	return nil
}

func saveImage(path string, imgFile io.ReadCloser, fileExt string) error {
	imageHandler := imageHandlerMap[fileExt]
	if imageHandler == nil {
		return fmt.Errorf("Missing config for image type %v ", fileExt)
	}

	return imageHandler(path, imgFile)
}

func scaleAndSaveWebp(path string, file io.ReadCloser) error {
	output, err := os.Create(path)
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

func convertJpegAndSaveWebp(path string, file io.ReadCloser) error {
	return convertImageAndSaveWebp(path, file, jpeg.Decode)
}

func convertPngAndSaveWebp(path string, file io.ReadCloser) error {
	return convertImageAndSaveWebp(path, file, png.Decode)
}

func convertImageAndSaveWebp(path string, file io.ReadCloser, decodeFunc DecodeFunc) error {
	img, err := decodeFunc(file)
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

	output, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create path for image %v ", err)
	}
	defer output.Close()

	if err := webp.Encode(output, scaledImg, options); err != nil {
		os.Remove(path) // remove the file created on failed write
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
