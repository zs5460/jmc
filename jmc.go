// Copyright 2021 zs. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jmc

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/zs5460/jm"
)

var (
	// K is an AES key that can be used for AES encryption and decryption operations. You should set an environment variable called "JMC_K" to change it.
	K  = "zs5460@gmail.com"
	re = regexp.MustCompile(`\${enc:([^}]+)}`)
)

var configParser = make(map[string]func(string, interface{}) error)

func init() {
	configParser[".json"] = LoadJSONConfig
	configParser[".xml"] = LoadXMLConfig

	p := os.Getenv("JMC_K")
	if len(p) == 16 || len(p) == 24 || len(p) == 32 {
		K = p
	}

}

// Encode ...
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

// Decode ...
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
	configFileType := path.Ext(fn)
	parser, exist := configParser[configFileType]
	if !exist {
		panic("file types that are not supported.")
	}
	err := parser(fn, v)
	if err != nil {
		panic("config file parse error.")
	}
}

// GetAppConfig returns the config path name for the executable that started
// the current process.
func GetAppConfig() string {
	app, err := os.Executable()
	if err != nil {
		panic(err)
	}
	if ext := filepath.Ext(app); ext != "" {
		app = strings.TrimRight(app, ext)
	}
	app += ".json"
	return app
}
