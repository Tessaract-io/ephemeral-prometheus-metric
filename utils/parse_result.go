package utils

import (
	"ephemeral-prometheus-metric/models"
	"strings"
)

// ParseResult function is used to parse the result of the command `df -k`.
// Extract the
func ParseResult(perLineResult []string, volumePath string) *models.StorageData {
	data := make(map[string]float64, 0)
	indexToColumnName := make(map[int]string, 0)
	stop := false
	for idx, line := range perLineResult {
		// split to match with the columns
		if idx == 0 {
			columnName := strings.Fields(line)
			// Linux : `Filesystem` `1K-blocks`   `Used` `Available`                            `Use%`   `Mounted` `on`
			// Darwin: `Filesystem` `1024-blocks` `Used` `Available` `Capacity` `iused` `ifree` `%iused` `Mounted` `on`
			for i, name := range columnName {
				data[name] = float64(0)
				indexToColumnName[i] = name
			}
			continue
		}
		items := strings.Fields(line)
		for i, item := range items {
			columnName := indexToColumnName[i]
			ok, floatItem := ConvertToFloat64(item)
			if ok {
				data[columnName] = floatItem
			} else {
				data[columnName] = 0
			}
			if columnName == "Mounted" && item == volumePath {
				// only capturing <volumePath> data
				stop = true
				break
			}
		}
		if stop {
			break
		}
	}
	return MapResult(data)
}
