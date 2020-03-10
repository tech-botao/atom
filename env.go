package lib

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/thoas/go-funk"
	"os"
	"path/filepath"
)

var _path = "/etc/go"

// 转换路径到abs
func SetEnvPath(path string) {
	path, err := filepath.Abs(path)
	if err != nil {
		panic( err)
	}
	_path = path
}

func GetEnvPath() string {
	return _path
}

func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

func GetEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	return env
}

// 是否微开发环境
func IsDev() bool {
	return GetEnv() == "dev"
}

// 加载不同Env文件
func LoadEnv(filenames ...string) {

	//logger.Info("[env] env folder", _path)

	filenames = filenamesOrDefault(filenames)
	configFiles := funk.Map(filenames, func(filename string) string {
		return GetEnvPath() + filename
	})

	err := godotenv.Load(configFiles.([]string)...)
	if err != nil {
		panic(err)
		//logger.Panic("config file:", filenames)
		//logger.Panic("read env file error: ", err)
	}
}

// 环境变量没有的话直接报错不犹豫
func EnvOrPanic(key string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		panic(errors.New("[env] empty env var for key="+ key))
	}
	return s
}
