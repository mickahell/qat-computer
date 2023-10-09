package utils

import (
	"encoding/json"
	"os"
	"qat-computer/logger"
)

func ToJSON(obj interface{}) string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.GetLogger().LogCritical("utils", "error with json serialization", err)
	}

	return string(res)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
