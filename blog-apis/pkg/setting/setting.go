package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunMode   string
	PageSize  int
	JwtSecret string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	DBType        string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBName        string
	DBTablePrefix string
)

func init() {
	var err error

	Cfg, err = ini.Load("F:\\go3\\blog-apis\\conf\\api.ini")
	if err != nil {
		log.Fatalf("无法加载“conf/api.ini”：%v", err)
	}

	LoadBase()
	LoadApp()
	LoadServer()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("无法获取“app”配置项：%v", err)
	}

	PageSize = section.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = section.Key("JWT_SECRET").MustString("123#456")
}

func LoadServer() {
	section, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("无法获取“server”配置项：%v", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadDatabase() {
	section, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("无法获取“database”配置项：%v", err)
	}

	DBType = section.Key("TYPE").MustString("mysql")
	DBUser = section.Key("USER").MustString("root")
	DBPassword = section.Key("PASSWORD").MustString("root")
	DBHost = section.Key("HOST").MustString("127.0.0.1:3306")
	DBName = section.Key("NAME").MustString("")
	DBTablePrefix = section.Key("TABLE_PREFIX").MustString("")
}
