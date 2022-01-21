package utils

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github/gin-react-admin/global"
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

func SaveUploadedFile(file *multipart.FileHeader, mk string, filename string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	eisdir, direr := PathExists(mk)      // 判断文件夹是否存在
	if direr != nil && eisdir == false { // 若不存在
		maker := os.MkdirAll(mk, os.ModePerm)
		if maker != nil {
			return  maker
		}
	}
	out, errors := os.Create(mk + "/" + filename)
	if errors != nil {
		return  errors
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
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

func CompileFileName (fileName string, cove int) (string, string) {
	fnfix := path.Ext(fileName)               // 获取文件名后缀
	fn := strings.TrimSuffix(fileName,fnfix)  //获取文件名
	ere := regexp.MustCompile("\\((\\d?)\\)") // 匹配括号
	_fName := regexp.MustCompile("(.*)\\(")
	machArre := ere.FindStringSubmatch(fn)
	newName := _fName.FindStringSubmatch(fn)
	if machArre == nil {
		return fn + "(1)" + fnfix , fnfix
	}
	fineNum, e2 := strconv.Atoi(machArre[1])
	if e2 != nil {
		return fn + "(1)"+ fnfix , fnfix
	}
	fineNum = cove
	newFileName := fmt.Sprintf("%v(%v)",newName[1], fineNum)
	return newFileName ,fnfix
}

func SaveStrFile (str string, mk string, name string) (string ,error) {
	dis, er := os.Create(mk + name)
	if er != nil {
		return "" ,er
	}
	defer dis.Close()
	eisdir, direr := PathExists(mk)      // 判断文件夹是否存在
	if direr != nil && eisdir == false { // 若不存在
		maker := os.MkdirAll(mk, os.ModePerm)
		if maker != nil {
			return "",maker
		}
	}
	_, err := dis.WriteString(str)
	return mk ,err
}

func bufferWrite(param string, filename string) (err error) {
	fileHandle, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriter(fileHandle)
	// 字节写入
	_ ,err = buf.Write([]byte("buffer Write : " + param))
	// 字符串写入
	_ ,err = buf.WriteString("buffer WriteString : " + param)
	// 将缓冲中的数据写入
	err = buf.Flush()
	return err
}