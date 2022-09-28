package config

import (
	"os"
	"path/filepath"
)

type IConfig struct {
	AppPath   string
	KeyPath   string
	TokenPath string
}

var userPath, _ = os.UserConfigDir();
var cominnekPath = filepath.Join(userPath, ".cominnek")

var Public = IConfig{
	KeyPath:   filepath.Join(cominnekPath, "key.bin"),
	TokenPath: filepath.Join(cominnekPath,"auth.bin"),
	AppPath:   cominnekPath,
}