package model

type File struct {
	FileName string `gorm:"comment:文件名字"`
	SavaPath string `gorm:"comment:文件的存储路径"`

}
