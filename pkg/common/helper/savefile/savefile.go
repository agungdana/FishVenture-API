package savefile

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

const (
	JPG  = "jpg"
	JPEG = "jpeg"
	PNG  = "png"
	GIF  = "gif"
	TIFF = "tiff"
	PSD  = "psd"
	PDF  = "pdf"
	EPS  = "eps"
	AI   = "ai"
	RAW  = "raw"
	WEBP = "webp"
)

var ImageExt = map[string]struct{}{
	JPG:  {},
	JPEG: {},
	PNG:  {},
	GIF:  {},
	TIFF: {},
	PSD:  {},
	EPS:  {},
	AI:   {},
	RAW:  {},
	WEBP: {},
}

var FileExt = map[string]struct{}{
	PDF: {},
}
