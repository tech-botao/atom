package atom

import (
	"github.com/joho/godotenv"
	"github.com/tech-botao/logger"
	"github.com/thoas/go-funk"
	"os"
	"path/filepath"
)

var _path = "/etc/go"

// 转换路径到abs
func SetEnvPath(path string) {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	_path = path
}

// 是否可以访问
func IsAccessible() {
	if f, err := os.Stat(_path); os.IsNotExist(err) || !f.IsDir() {
		logger.Panic("[env] path is not accessible", _path)
	}
}

// 取得并检查配置目录
func GetEnvPath() string {
	IsAccessible()
	return _path
}

// 配置文件
func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{"env"}
	}
	return filenames
}

// 取得环境变量, prod, test, dev, 默认dev
func GetAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	return env
}

// 是否微开发环境
func IsDev() bool {
	return GetAppEnv() == "dev" || GetAppEnv() == "test"
}

// 加载不同Env文件
func LoadEnv(filenames ...string) {

	logger.Info("[env] env folder", _path)
	filenames = filenamesOrDefault(filenames)
	configFiles := funk.Map(filenames, func(filename string) string {
		return GetEnvPath() + "/" + filename
	})

	err := godotenv.Load(configFiles.([]string)...)
	if err != nil {
		logger.Info("config file:", filenames)
		logger.Panic("read env file error: ", err)
	}
}

// 环境变量没有的话直接报错不犹豫
func EnvOrPanic(key string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		logger.Panic("[env] empty env var for key=" + key, nil)
	}
	return s
}
