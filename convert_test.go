package lib

import (
	"github.com/k0kubun/pp"
	"io"
	"reflect"
	"strings"
	"testing"
)

func generatorCase(t string) (io.Reader, interface{}) {
	switch t {
	case "structure":
		reader := strings.NewReader(`{"float":1234.1,"int":123,"string":"abc","structure":{"int":100,"string":"s0s"}}`)
		type structure struct{
			Int int	`json:"int"`
			String string `json:"string"`
			Float float64 `json:"float"`
			Structure struct{
				Int int `json:"int"`
				String string `json:"string"`
			}
		}
		return reader, structure{}
	case "array":
		reader := strings.NewReader(`[{"float":1234.1,"int":123,"string":"abc"}]`)
		type structure struct{
			Int int
			String string
			Float float64
		}
		return reader, []structure{}
	default:
		return nil, nil
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		caseType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"struct_ok", args{"structure"}, false},
		{"array_ok", args{"array"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, structure := generatorCase(tt.args.caseType)
			if err := DecodeFromReader(reader, &structure); (err != nil) != tt.wantErr {
				t.Errorf("DecodeFromReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			pp.Println(structure)
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		caseType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"struct_ok", args{"structure"}, false},
		{"array_ok", args{"array"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, structure := generatorCase(tt.args.caseType)
			expect := ReaderToString(reader)
			_ = DecodeFromReader(strings.NewReader(expect), &structure)
			got, err := Encode(structure)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}


			if got != expect {
				t.Errorf("Encode() got = %v, want %v", got, expect)
			}
		})
	}
}

func TestEncodeToReader(t *testing.T) {
	type args struct {
		A string `json:"a,omitempty"`
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"empty", args{}, `{}`, false},
		{"value", args{"chilema"}, `{"a":"chilema"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeToReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeToReader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReaderToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ok", "{}"},
		{"ok", `{"a":"1234"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.want)
			if got := ReaderToString(reader); got != tt.want {
				t.Errorf("ReaderToString() = %v, want %v", got, tt.want)
			}
		})
	}
}