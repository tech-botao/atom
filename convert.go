package lib

import (
	"bytes"
	"encoding/json"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)


// DecodeFromByte []byte -> struct
func DecodeFromByte(b []byte, v interface{}) error {
	buff := bytes.NewReader(b)
	return json.NewDecoder(buff).Decode(&v)
}

// DecodeFromFile jsonFile -> struct
func DecodeFromFile(filename string, v interface{}) error {
	f, err := ioutil.ReadFile(filename)
	pp.Println(err)
	if err != nil {
		return errors.WithMessage(err, filename + " open error")
	}

	return json.NewDecoder(ByteToReader(f)).Decode(&v)
}

// DecodeFromReader json -> struct
func DecodeFromReader(reader io.Reader, v interface{}) error {
	if reader == nil {
		return errors.Errorf("DecodeFromReader() reader=nil")
	}
	return json.NewDecoder(reader).Decode(&v)
}

// 测试用的函数，输出内容
func ReaderDump(reader io.Reader) error {
	s := ReaderToString(reader)
	pp.Println(s)
	return nil
}

// DecodeFromReader string -> json
func DecodeFromString(str string, v interface{}) error {
	return json.NewDecoder(strings.NewReader(str)).Decode(&v)
}

// Encode struct -> json
func Encode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

// EncodeToBytes struct -> []byte
func EncodeToBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func EncodeToReader(v interface{}) (io.Reader, error) {
	s, err := Encode(v)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(s),nil
}

// io.Reader -> string
func ReaderToString(reader io.Reader) string {
	if reader == nil {
		return ""
	}
	buff := new (bytes.Buffer)
	_, _ = buff.ReadFrom(reader)
	return buff.String()
}

func ByteToReader(data []byte) io.Reader {
	return bytes.NewReader(data)
}

// ======================= time =======================

func Time() string {
	return time.Now().Format("15:04:05")
}

func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Date() string {
	return time.Now().Format("2006-01-02")
}

func HHMMSS2Time(s string) time.Time {
	t, err := time.Parse("15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}

func Str2Time(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}

// ========= math =============
// utils
func StrToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

