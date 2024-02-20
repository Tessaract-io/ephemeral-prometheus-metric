package tests

import (
	"ephemeral-prometheus-metric/utils"
	"testing"
)

func TestDefaultResult(t *testing.T) {
	result := utils.DefaultResult()
	if result.TotalCapacityBytes != 0 {
		t.Fatalf("TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != 0 {
		t.Fatalf("UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != 0 {
		t.Fatalf("RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != 0 {
		t.Fatalf("UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != 0 {
		t.Fatalf("RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}
