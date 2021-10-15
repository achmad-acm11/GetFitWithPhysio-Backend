package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadPhoto(req *http.Request, nameFile string, formName string, dirPath string) string {
	file, fileHeader, err := req.FormFile(formName)
	HandleError(err)
	defer file.Close()

	dir, _ := os.Getwd()

	filename := fmt.Sprintf("%s%s", nameFile, filepath.Ext(fileHeader.Filename))
	pathFile := fmt.Sprintf("team_photos/%s", filename)

	path := filepath.Join(dir, dirPath, filename)
	target, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer target.Close()

	_, err = io.Copy(target, file)
	HandleError(err)

	return pathFile
}
