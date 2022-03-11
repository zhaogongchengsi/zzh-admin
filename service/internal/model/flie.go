package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName string `gorm:"comment:文件名字" json:"file_name"`
	SavaPath string `gorm:"comment:文件的存储路径" json:"sava_path"`
	FileExt string `gorm:"comment:文件的后缀" json:"file_ext"`
	FileBroadType string `gorm:"comment:文件广泛的类型" json:"file_broad_type"`
	FileSpecificType string `gorm:"comment:文件具体的类型" json:"file_specific_type"`
	IsHash bool `gorm:"comment:文件名是否被hash" json:"is_hash"`
	Overwrite bool `gorm:"comment:是否覆盖" json:"overwrite"`
	CoverageTimes int `gorm:"comment:覆盖次数" json:"coverage_times"`
}
