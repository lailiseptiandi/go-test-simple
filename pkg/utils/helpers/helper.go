package helpers

import (
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

func UploadImageFile(file multipart.File, header *multipart.FileHeader) string {

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "/" + os.Getenv("STORAGE_PATH") + "/" + filename

	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	path := fmt.Sprintf("%v/%v", os.Getenv("STORAGE_PATH"), filename)
	err = imaging.Save(src, path)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return filePath
}
