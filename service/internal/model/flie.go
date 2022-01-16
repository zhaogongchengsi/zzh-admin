package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName string `gorm:"comment:文件名字"`
	SavaPath string `gorm:"comment:文件的存储路径"`
	FileExt string `gorm:"comment:文件的后缀"`
	FileBroadType string `gorm:"comment:文件广泛的类型"`
	FileSpecificType string `gorm:"comment:文件具体的类型"`
	IsHash bool `gorm:"comment:文件名是否被hash"`
}
