package utils

import "ephemeral-prometheus-metric/models"

func MapResult(data map[string]float64) *models.StorageData {
	totalCapacityBytes, ok := data["1K-blocks"]
	if !ok {
		// Darwin
		totalCapacityBytes, ok = data["1024-blocks"]
		if !ok {
			logger.Printf("[ERROR  ] Cannot find the total capacity bytes\n")
			return DefaultResult()
		}
	}
	totalCapacityBytes = totalCapacityBytes * 1024
	remainingBytes, ok := data["Available"]
	if !ok {
		logger.Printf("[ERROR  ] Cannot find the remaining bytes\n")
		return DefaultResult()
	}
	remainingBytes = remainingBytes * 1024

	usageBytes := totalCapacityBytes - remainingBytes
	return &models.StorageData{
		TotalCapacityBytes: totalCapacityBytes,
		RemainingBytes:     remainingBytes,
		RemainingPercent:   remainingBytes / totalCapacityBytes * 100,
		UsageBytes:         usageBytes,
		UsagePercent:       usageBytes / totalCapacityBytes * 100,
	}
}
