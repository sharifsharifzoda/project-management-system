package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GenFilenameWithDir(filename string) (string, error) {
	validExtensions := []string{".jpg", ".jpeg", ".png"}

	fileExt := strings.ToLower(filepath.Ext(filename))

	for _, ext := range validExtensions {
		if ext == fileExt {
			timeSign := fmt.Sprintf("%d", time.Now().UnixNano())
			filePath := fmt.Sprintf("%s_%s", timeSign, filename)
			filePath = strings.Replace(filePath, " ", "", -1)

			return filePath, nil
		}
	}

	return "", errors.New("unsupported media type")
}
