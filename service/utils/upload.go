package utils

import (
	"fmt"
	"github/gin-react-admin/global"
	"io"
	"mime/multipart"
	"os"
	"time"
)



func MkDir (fileName string) string {
	var upPath string = global.FileService.UploadPath
	now := time.Now() //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	return fmt.Sprintf("%s/%v/%v",upPath, year,month)
}

func SaveUploadedFile(file *multipart.FileHeader, mk string ,dst string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "" , err
	}
	defer src.Close()
	eisdir, direr := PathExists(mk)      // 判断文件夹是否存在
	if direr != nil && eisdir == false { // 若不存在
		maker := os.MkdirAll(mk, os.ModePerm)
		if maker != nil {
			return "", maker
		}
	}
	var filePath string = mk + "/" + dst
	out, errors := os.Create(filePath)
	if errors != nil {
		return "", errors
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return filePath, err
}

func PathExists (path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true , nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return  false, err
}