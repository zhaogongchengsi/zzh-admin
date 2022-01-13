package main

import (
	"fmt"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/routers"
	"github/gin-react-admin/pkg/setting"
	"log"
	"net/http"
	"time"

	"github/gin-react-admin/global"

	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func init() {
	err := setupSettings() // 初始化读取配置
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	RunInputMsg()
}

// @title gin-react-admin API 文档
// @version 0.0.1
// @description gin-react-admin 接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name r-token
// @BasePath /
func main() {
	gin.SetMode(global.ServerSetting.RenMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()

	if err != nil {
		return
	}
}

func setupSettings() error {

	newSetting, err := setting.NewSetting()
	cosSetting, err := setting.NewSetOsConfig()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Captcha", &global.Captcha)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("FileService", &global.FileService)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.Jwt)
	if err != nil {
		return err
	}
	err = cosSetting.ReadSection("CosConfig", &global.Cos)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func RunInputMsg() {
	fmt.Printf(`
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1:%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:3000
	默认静态文件服务运行在:http://127.0.0.1:%s/%v
	运行 swag init 生成最新文档
`, global.ServerSetting.HttpPort,global.ServerSetting.HttpPort, "assets")
}
