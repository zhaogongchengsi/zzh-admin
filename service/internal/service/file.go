package service

import (
	"fmt"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/utils"
	"mime/multipart"
	"path"
	"strings"
)

func FileSave (file *multipart.FileHeader, isHash bool) (model.File ,error) {
	fileType := strings.Split(file.Header.Get("Content-Type"), "/") // 文件类型
	_fileMode := model.File{
		FileBroadType:fileType[0],
		FileSpecificType: fileType[1],
		IsHash: isHash,
		Overwrite: false,
	}
	err := ConfirmFileName(file.Filename, &_fileMode)
	upper := utils.SaveUploadedFile(file, _fileMode.SavaPath , _fileMode.FileName + _fileMode.FileExt)
	if upper != nil {
		return _fileMode, upper
	}
	err = global.DBEngine.Create(&_fileMode).Error
	return _fileMode, err
}

func FirstFile(fileName string, fileExr string) ([]model.File ,error)  {
	var files []model.File
 	e := global.DBEngine.Model(&model.File{}).Where("`file_name` LIKE ? AND `coverage_times` > ? AND file_ext = ?", "%" + fileName + "%", 0, fileExr).Find(&files).Error
	return files, e
}


func SaveString (str string, fame string, Ashish bool) (*model.File, error) {
	_fileMode := model.File{
		FileBroadType:    "text",
		FileSpecificType: "text",
		IsHash:           Ashish,
		Overwrite:        false,
	}
	cert := ConfirmFileName(fame, &_fileMode)
	if cert != nil {
		return &model.File{}, cert
	}
	newPath , serer := utils.SaveStrFile(str, _fileMode.SavaPath , _fileMode.FileName + _fileMode.FileExt)
	if serer != nil {
		return &_fileMode, serer
	}
	_fileMode.SavaPath = newPath
	err := global.DBEngine.Create(&_fileMode).Error
	return &_fileMode, err
}

// ConfirmFileName 确认文件上是否覆盖（文件重复次数） 文件名 文件后缀
func ConfirmFileName (file string, filo *model.File) error {
	var fileName, fileExt string
	var newCove int = 1
	fileExt = path.Ext(file)               // 获取文件名后缀
	fileName = strings.TrimSuffix(file, fileExt)
	filePath := utils.MkDir(fileName)
	if filo.Overwrite == false {
		files, er := FirstFile(fileName, fileExt) // 确认上传的文件是否已经存在
		if er != nil {
			return er
		}
		if filesSize := len(files); filesSize > 0 {
			oldFile := files[filesSize - 1]
			newCove = oldFile.CoverageTimes + 1
			if filo.IsHash == true {
				fileName  =  utils.HashFileName(fileName)
			}
			fileName = fileName+fmt.Sprintf("(%v)", newCove)
		}
	}
	filo.FileName = fileName
	filo.FileExt = fileExt
	filo.CoverageTimes = newCove
	filo.SavaPath = filePath + "/"
	return nil
}

