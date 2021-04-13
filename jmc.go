package jmc

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/zs5460/jm"
)

var (
	K  = "zs5460@gmail.com"
	re = regexp.MustCompile(`\${enc:([^}]+)}`)
)

var configParser = make(map[string]func(string, interface{}) error, 0)

func init() {
	configParser[".json"] = LoadJSONConfig
	configParser[".xml"] = LoadXMLConfig
}

func Encode(s string) string {
	if !strings.Contains(s, "${enc:") {
		return s
	}
	result := re.FindAllStringSubmatch(s, -1)
	for _, sm := range result {
		enc, _ := jm.EncryptString(sm[1], K)
		s = strings.Replace(s, sm[1], enc, -1)
	}
	return s
}

func Decode(s string) (string, error) {
	if !strings.Contains(s, "${enc:") {
		return s, nil
	}
	result := re.FindAllStringSubmatch(s, -1)
	for _, sm := range result {
		dec, err := jm.DecryptString(sm[1], K)
		if err != nil {
			return "", err
		}
		s = strings.Replace(s, sm[0], dec, -1)
	}
	return s, nil
}

// LoadJSONConfig load config from json file.
func LoadJSONConfig(fn string, v interface{}) error {
	content, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	decoded, err := Decode(string(content))
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(decoded), v)
	return err
}

// LoadXMLConfig load config from xml file.
func LoadXMLConfig(fn string, v interface{}) error {
	content, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	decoded, err := Decode(string(content))
	if err != nil {
		return err
	}
	err = xml.Unmarshal([]byte(decoded), v)
	return err
}

// MustLoadConfig load config or panic.
func MustLoadConfig(fn string, v interface{}) {
	configtype := path.Ext(fn)
	parser, exist := configParser[configtype]
	if !exist {
		panic("unsupported config file.")
	}
	err := parser(fn, v)
	if err != nil {
		panic("config file parse error.")
	}
}
