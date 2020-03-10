package lib

import (
	"fmt"
	"github.com/tech-botao/logger"
	"os"
	"reflect"
	"testing"
)

func TestSetEnvPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{".", args{"."}},
		//{"~", args{"~/"}}, // 这个符号用不了
		{"../", args{"../"}},
		{"log", args{"log"}},
		{"/etc", args{"/etc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetEnvPath(tt.args.path)
			fmt.Println(tt.args.path, "=>", _path)
		})
	}
}

func TestEnvOrPanic(t *testing.T) {
	t.Skipf("include EnvLoad()")
}

func TestGetAppEnv(t *testing.T) {
	t.Skipf("include EnvLoad()")
}

func TestGetEnvPath(t *testing.T) {
	t.Skipf("include EnvLoad()")
}

func TestIsAccessible(t *testing.T) {
	t.Skipf("include EnvLoad()")
}

func TestIsDev(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"dev", true},
		{"test", true},
		{"prod", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("APP_ENV", tt.name)
			if got := IsDev(); got != tt.want {
				t.Errorf("IsDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadEnv(t *testing.T) {
	type args struct {
		filenames []string
	}
	SetEnvPath(".")
	tests := []struct {
		name string
		args args
	}{
		{"atom.env", args{filenames:filenamesOrDefault(nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadEnv(tt.args.filenames...)
			logger.Info("APP_ENV", GetAppEnv())
			logger.Info("LOG_FILE", EnvOrPanic("LOG_FILE"))
		})
	}
}


func Test_filenamesOrDefault(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"default", args{[]string{}}, []string{"env"}},
		{"one file", args{[]string{"atom.env"}}, []string{"atom.env"}},
		{"two file", args{[]string{"atom.env", "atom2.env"}}, []string{"atom.env", "atom2.env"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filenamesOrDefault(tt.args.filenames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filenamesOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}