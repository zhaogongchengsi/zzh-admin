package setting

import "time"

type ServerSettingS struct {
	RenMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DafaultPageSize int
	MaxPageSize     int
	LogSavepath     string
	LogFilename     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type Captcha struct {
	ImgWidth  int
	ImgHeight int
	KeyLong   int
	MaxSkew   float64
	DotCount  int
}

type JWT struct {
	Issuer     string
	SigningKey string
	ExpiresAt  int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
