package commands

import (
	"ephemeral-prometheus-metric/models"
	"ephemeral-prometheus-metric/utils"
)

// GetStorageData function is used to get the <path> storage data.
// Output is following the format of StorageData struct.
func GetStorageData(path string) *models.StorageData {
	errMsg, cmdResult := RunCommand("df", []string{"-k"})
	if errMsg != "" {
		logger.Printf("[ERROR  ] %s\n", errMsg)
		return utils.DefaultResult()
	}
	return utils.ParseResult(cmdResult, path)
}
