package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type HttpConfig struct {
    Url     string
    Port    string
}

type ConfigT struct {
    Http        HttpConfig
}

var Config = ConfigT{
    Http: HttpConfig{
        Url:    "",
        Port:   "3001",
    },
}

func InitConfig() {
    ex, err := os.Executable()
    if nil != err {
        panic(err)
    }
    expath := filepath.Dir(ex)
    configfile := expath + "/.config.json"

    dat, err := os.ReadFile(configfile)
    if nil != err {
        configdat, _ := json.MarshalIndent(Config, "", "  ")
        os.WriteFile(configfile, configdat, 0644)
    } else {
        json.Unmarshal(dat, &Config)
    }
}
