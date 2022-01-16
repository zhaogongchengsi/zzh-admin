package service

import (
	"fmt"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/utils"
	"mime/multipart"
)

func FileSave (file *multipart.FileHeader, isHash bool) (model.File ,error) {
	filePath := utils.MkDir(file.Filename)
	// 上传文件到指定的 dst
	_file, _ := FirstFile(file.Filename)

	fmt.Printf("文件已存在 %v", _file)

	fileModel, upper := utils.SaveUploadedFile(file, filePath, isHash)
	if upper != nil {
		return fileModel, upper
	}



	err := global.DBEngine.Create(&fileModel).Error
	return fileModel, err
}

func FirstFile(fileName string) (model.File ,error)  {
	var file model.File
	e := global.DBEngine.Where("file_name = ?", fileName).First(&file).Error
	return file, e
}