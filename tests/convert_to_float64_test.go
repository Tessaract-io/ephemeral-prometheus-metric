package tests

import (
	"ephemeral-prometheus-metric/utils"
	"testing"
)

func TestConvertToFloat64(t *testing.T) {
	data := "123"
	ok, result := utils.ConvertToFloat64(data)
	if ok != true {
		t.Fatalf("Test ConvertToFloat64: ok is not correct, value is %t", ok)
	}
	if result != 123 {
		t.Fatalf("Test ConvertToFloat64: result is not correct, value is %f", result)
	}
}

func TestConvertToFloat64Fail(t *testing.T) {
	data := "123a"
	ok, result := utils.ConvertToFloat64(data)
	if ok != false {
		t.Fatalf("Test ConvertToFloat64Fail: ok is not correct, value is %t", ok)
	}
	if result != 0 {
		t.Fatalf("Test ConvertToFloat64Fail: result is not correct, value is %f", result)
	}
}
