package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/payfazz/chrome-remote-debug/config"
)

func init() {
	once.Do(func() {
		config.SetConfig(&env{})
	})
}

type env struct{}

var e *env
var once sync.Once

//  configuration keys
const (
	chromeDownloadPath = "CHROME_DOWNLOAD_PATH"
)

func (e *env) ChromeDownloadPath() string {
	return getStringOrDefault(chromeDownloadPath, "/home/chrome/Downloads")

}

func getMapOrDefault(key string, def map[string]interface{}) map[string]interface{} {
	b, err := json.Marshal(def)
	if err != nil {
		return def
	}

	v := getEnvOrDefault(key, string(b))

	var fm map[string]interface{}
	err = json.Unmarshal([]byte(v), &fm)
	if err != nil {
		return def
	}

	return fm
}

// Helper methods
func getStringOrDefault(key, def string) string {
	return getEnvOrDefault(key, def)
}

func getIntOrDefault(key string, def int) int {
	v := getEnvOrDefault(key, fmt.Sprint(def))
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return def
	}
	return int(i)
}

func getBooleanOrDefault(key string, def bool) bool {
	v := getEnvOrDefault(key, fmt.Sprint(def))
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}

func getTime(key string) time.Time {
	v := getEnvOrDefault(key, "1970-01-01T00:00:00+07:00")
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		now := time.Now()
		return time.Date(1970, 1, 1, 0, 0, 0, 0, now.Location())
	}
	return t
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
