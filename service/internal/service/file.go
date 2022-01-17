package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/utils"
	"mime/multipart"
	"strings"
)

func FileSave (file *multipart.FileHeader, isHash bool) (model.File ,error) {
	filePath := utils.MkDir(file.Filename)
	fileName, fileExt := utils.CompileFileName(file.Filename)
	fileType := strings.Split(file.Header.Get("Content-Type"), "/") // 文件类型
	// 上传文件到指定的 dst
	_file, er := FirstFile(file.Filename)
	if er == nil {
		fileName, _ = utils.CompileFileName(_file.FileName)
	} else {

	}
	if isHash == true {
		fileName  =  utils.HashFileName(fileName)
	}
	_fileMode := model.File{
		FileExt: fileExt ,
		FileName: fileName + fileExt,
		FileBroadType:fileType[0],
		FileSpecificType: fileType[1],
		IsHash: isHash,
	}

	fileModel, upper := utils.SaveUploadedFile(file, filePath, &_fileMode)
	if upper != nil {
		return fileModel, upper
	}
	err := global.DBEngine.Create(&_fileMode).Error
	return fileModel, err

	//fileModel, upper := utils.SaveUploadedFile(file, filePath, &_fileMode)
	//if upper != nil {
	//	return fileModel, upper
	//}
	//err := global.DBEngine.Model(&_file).Updates(&fileModel).Error

	return fileModel, err
}

func FirstFile(fileName string) (model.File ,error)  {
	var file model.File
	e := global.DBEngine.Where("file_name = ?", fileName).First(&file).Error
	return file, e
}