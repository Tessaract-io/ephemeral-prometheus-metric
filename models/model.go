package models

type StorageData struct {
	TotalCapacityBytes float64
	RemainingBytes     float64
	RemainingPercent   float64
	UsagePercent       float64
	UsageBytes         float64
}
