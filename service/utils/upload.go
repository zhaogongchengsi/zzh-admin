package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"io"
	"mime/multipart"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)



func MkDir (fileName string) string {
	var upPath string = global.FileService.UploadPath
	now := time.Now() //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	return fmt.Sprintf("%s/%v/%v",upPath, year,month)
}

func SaveUploadedFile(file *multipart.FileHeader, mk string, Ashish bool) (model.File, error) {
	src, err := file.Open()
	if err != nil {
		return model.File{} , err
	}
	CompileFileName(file.Filename)
	dst := file.Filename // 文件名
	ext := path.Ext(dst) // 文件后缀
	fileType := strings.Split(file.Header.Get("Content-Type"), "/") // 文件类型
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
			return model.File{}, maker
		}
	}



	var filePath string = mk + "/" + filename
	out, errors := os.Create(filePath)
	if errors != nil {
		return model.File{}, errors
	}
	defer out.Close()
	_, err = io.Copy(out, src)

	fileObj := model.File{
		FileName: filename,
		FileBroadType:fileType[0],
		FileSpecificType: fileType[1],
		FileExt: ext,
		SavaPath: filePath,
		IsHash: Ashish,
	}

	return fileObj, err
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

// 将文件名Hash化 防止文件上传攻击
func HashFileName(filename string) string {
	key := global.FileService.HashKey + filename
	sha1h := sha1.New()
	shame := sha1h.Sum([]byte(key))
	return hex.EncodeToString(shame)
}

func CompileFileName (fileName string) (string, string) {
	fnfix := path.Ext(fileName)               // 获取文件名后缀
	fn := strings.TrimSuffix(fileName,fnfix)  //获取文件名
	ere := regexp.MustCompile("\\((\\d?)\\)") // 匹配括号
	_fName := regexp.MustCompile("(.*)\\(")
	//nReg := regexp.MustCompile("-?[1-9]\\d*") // 匹配数字
	//zfretg := regexp.MustCompile("（(.*?)）") // 匹配中文括号
	machArre := ere.FindStringSubmatch(fn)
	newName := _fName.FindStringSubmatch(fn)
	//num := nReg.FindStringSubmatch(machArre[1])
	//machArrz := zfretg.FindStringSubmatch(fn)
	if len(machArre) < 0 {
		return fn + "(1)", fnfix
	}

	fineNum, e2 := strconv.Atoi(machArre[1])
	if e2 != nil {
		return fn + "(1)", fnfix
	}
	fineNum++
	newFileName := fmt.Sprintf("%v(%v)",newName[1], fineNum)

	return newFileName ,fnfix
}