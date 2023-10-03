package utils

import (
	"cat-computer/logger"
	"encoding/json"
)

func ToJSON(obj interface{}) string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.GetLogger().LogCritical("utils", "error with json serialization", err)
	}

	return string(res)
}
