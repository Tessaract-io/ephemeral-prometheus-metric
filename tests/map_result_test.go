package tests

import (
	"ephemeral-prometheus-metric/utils"
	"testing"
)

func TestMapResultLinux(t *testing.T) {
	data := map[string]float64{
		"1K-blocks": 52416492,
		"Used":      23709872,
		"Available": 28706620,
	}
	result := utils.MapResult(data)
	if result.TotalCapacityBytes != 52416492*1024 {
		t.Fatalf("Test Linux: TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != 23709872*1024 {
		t.Fatalf("Test Linux: UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != 28706620*1024 {
		t.Fatalf("Test Linux: RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != result.UsageBytes/result.TotalCapacityBytes*100 {
		t.Fatalf("Test Linux: UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != result.RemainingBytes/result.TotalCapacityBytes*100 {
		t.Fatalf("Test Linux: RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}

func TestMapResultDarwin(t *testing.T) {
	data := map[string]float64{
		"1024-blocks": 239362496,
		"Used":        9897844,
		"Available":   8652156,
	}
	result := utils.MapResult(data)
	totalCapacityBytes := data["1024-blocks"] * 1024
	remainingBytes := data["Available"] * 1024
	usageBytes := totalCapacityBytes - remainingBytes
	if result.TotalCapacityBytes != totalCapacityBytes {
		t.Fatalf("Test Darwin: TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != usageBytes {
		t.Fatalf("Test Darwin: UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != remainingBytes {
		t.Fatalf("Test Darwin: RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != usageBytes/totalCapacityBytes*100 {
		t.Fatalf("Test Darwin: UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != remainingBytes/totalCapacityBytes*100 {
		t.Fatalf("Test Darwin: RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}

func TestMapResultError(t *testing.T) {
	data := map[string]float64{
		"1K-blocks": 52416492,
		"Used":      23709872,
		"-":         28706620,
	}
	result := utils.MapResult(data)
	if result.TotalCapacityBytes != 0 {
		t.Fatalf("Test Error: TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != 0 {
		t.Fatalf("Test Error: UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != 0 {
		t.Fatalf("Test Error: RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != 0 {
		t.Fatalf("Test Error: UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != 0 {
		t.Fatalf("Test Error: RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}
