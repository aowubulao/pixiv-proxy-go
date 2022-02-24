package utils

import (
	"os"
	"strings"
)

func GetArg(key string) string {
	for _, v := range os.Args {
		if strings.HasPrefix(v, key) {
			return strings.Replace(v, key+"=", "", 1)
		}
	}
	return ""
}

func GetArgWithDefault(key string, defaultVal string) string {
	for _, v := range os.Args {
		if strings.HasPrefix(v, key) {
			return strings.Replace(v, key+"=", "", 1)
		}
	}
	return defaultVal
}
