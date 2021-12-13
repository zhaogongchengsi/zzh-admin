package model

import (
	"fmt"
	"github/gin-react-admin/global"
	"github/gin-react-admin/pkg/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	Id          uint32 `gorm:"primary_key" json:"id"`
	CreateBy    string `josn:"created_by"`
	ModiefiedBy string `josn:"modified_by"`
	CreateOn    uint32 `json:"create_on"`
	ModifiedOn  uint32 `json:"modified_on"`
	DeletedOn   uint32 `json:"deleted_on"`
	IsDel       uint8  `json:"is_del"`
}

// 链接数据库
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       s,     // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,                        // 使用单数表名，启用该选项后，`User` 表将是`user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 不适用物理外键
	})

	if err != nil {
		return nil, err
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能

	sqlDB, sqlerr := db.DB()

	if sqlerr != nil {
		return nil, sqlerr
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func InitMenu(m interface{}) error {
	err := global.DBEngine.AutoMigrate(&m)
	if err != nil {
		return err
	}
	return nil
}
