package utils

import "ephemeral-prometheus-metric/models"

func DefaultResult() *models.StorageData {
	return &models.StorageData{
		TotalCapacityBytes: 0,
		RemainingBytes:     0,
		RemainingPercent:   0,
		UsagePercent:       0,
		UsageBytes:         0,
	}
}
