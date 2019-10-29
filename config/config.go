package config

import (
	"sync"
)

// Config , contains application configuration
type config interface {
	ChromeDownloadPath() string
}

var cfg config
var once sync.Once
var isSet bool

// SetConfig sets active config to c
// this method should be called by config implementor
// and can only be called once
func SetConfig(c config) {
	if isSet {
		panic("config is set more than once")
	}
	once.Do(func() {
		cfg = c
		isSet = true
	})
}

// ChromeDownloadPath chrome download path
func ChromeDownloadPath() string {
	return cfg.ChromeDownloadPath()
}
