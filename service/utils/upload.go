package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github/gin-react-admin/global"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"
)



func MkDir (fileName string) string {
	var upPath string = global.FileService.UploadPath
	now := time.Now() //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	return fmt.Sprintf("%s/%v/%v",upPath, year,month)
}

func SaveUploadedFile(file *multipart.FileHeader, mk string, Ashish bool) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "" , err
	}
	dst := file.Filename // 文件名
	ext := path.Ext(dst) // 文件后缀
	var filename string
	if Ashish == true {
		filename  =  HashFileName(dst) + ext
	} else {
		filename  =  dst
	}

	defer src.Close()
	eisdir, direr := PathExists(mk)      // 判断文件夹是否存在
	if direr != nil && eisdir == false { // 若不存在
		maker := os.MkdirAll(mk, os.ModePerm)
		if maker != nil {
			return "", maker
		}
	}

	var filePath string = mk + "/" + filename
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

// HashFileName 将文件名Hash化 防止文件上传攻击
func HashFileName(filename string) string {
	key := global.FileService.HashKey + filename
	sha1h := sha1.New()
	shame := sha1h.Sum([]byte(key))
	return hex.EncodeToString(shame)
}