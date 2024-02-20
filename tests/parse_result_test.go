package tests

import (
	"ephemeral-prometheus-metric/utils"
	"testing"
)

func TestParseResultLinux(t *testing.T) {
	data := []string{
		"Filesystem     1K-blocks     Used Available Use% Mounted on",
		"overlay         52416492 23709872  28706620  46% /",
		"tmpfs              65536        0     65536   0% /dev",
		"tmpfs           16120444        0  16120444   0% /sys/fs/cgroup",
		"/dev/nvme0n1p1  52416492 23709872  28706620  46% /etc/hosts",
		"shm                65536       36     65500   1% /dev/shm",
		"tmpfs           16120444       12  16120432   1% /run/secrets/kubernetes.io/serviceaccount",
		"tmpfs           16120444        0  16120444   0% /proc/acpi",
		"tmpfs           16120444        0  16120444   0% /sys/firmware",
	}
	result := utils.ParseResult(data, "/")
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

func TestParseResultDarwin(t *testing.T) {
	data := []string{
		"Filesystem     1024-blocks      Used Available Capacity iused    ifree %iused  Mounted on",
		"/dev/disk3s1s1   239362496   9897844   8652156    54%  393025 86521560    0%   /",
		"devfs                  200       200         0   100%     692         0  100%   /dev",
		"/dev/disk3s6     239362496   6292508  10864896    37%       6 108648960    0%   /System/Volumes/VM",
		"/dev/disk3s2     239362496   5839260  10864896    35%    1097 108648960    0%   /System/Volumes/Preboot",
		"/dev/disk3s4     239362496     85188  10864896     1%      48 108648960    0%   /System/Volumes/Update",
		"/dev/disk1s2        512000      6164    492144     2%       1   4921440    0%   /System/Volumes/xarts",
		"/dev/disk1s1        512000      6236    492144     2%      28   4921440    0%   /System/Volumes/iSCPreboot",
		"/dev/disk1s3        512000      2596    492144     1%      50   4921440    0%   /System/Volumes/Hardware",
		"/dev/disk3s5     239362496 205348512  10864896    95% 1842380 108648960    2%   /System/Volumes/Data",
		"map auto_home            0         0         0   100%       0         0     -   /System/Volumes/Data/home",
	}
	result := utils.ParseResult(data, "/")
	totalCapacityBytes := result.TotalCapacityBytes
	remainingBytes := result.RemainingBytes
	usageBytes := totalCapacityBytes - remainingBytes
	if result.TotalCapacityBytes != 239362496*1024 {
		t.Fatalf("Test Darwin: TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != usageBytes {
		t.Fatalf("Test Darwin: UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != 8652156*1024 {
		t.Fatalf("Test Darwin: RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != usageBytes/totalCapacityBytes*100 {
		t.Fatalf("Test Darwin: UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != remainingBytes/totalCapacityBytes*100 {
		t.Fatalf("Test Darwin: RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}

func TestParseResultLinuxLogic(t *testing.T) {
	data := []string{
		"Filesystem     1K-blocks     Used Available Use% Mounted on",
		"overlay          1000000   700000    300000  70% /",
		"tmpfs              65536        0     65536   0% /dev",
		"tmpfs           16120444        0  16120444   0% /sys/fs/cgroup",
		"/dev/nvme0n1p1  52416492 23709872  28706620  46% /etc/hosts",
		"shm                65536       36     65500   1% /dev/shm",
		"tmpfs           16120444       12  16120432   1% /run/secrets/kubernetes.io/serviceaccount",
		"tmpfs           16120444        0  16120444   0% /proc/acpi",
		"tmpfs           16120444        0  16120444   0% /sys/firmware",
	}
	result := utils.ParseResult(data, "/")
	if result.TotalCapacityBytes != 1000000*1024 {
		t.Fatalf("Test Linux: TotalCapacityBytes is not correct, value is %f", result.TotalCapacityBytes)
	}
	if result.UsageBytes != 700000*1024 {
		t.Fatalf("Test Linux: UsageBytes is not correct, value is %f", result.UsageBytes)
	}
	if result.RemainingBytes != 300000*1024 {
		t.Fatalf("Test Linux: RemainingBytes is not correct, value is %f", result.RemainingBytes)
	}
	if result.UsagePercent != 70 {
		t.Fatalf("Test Linux: UsagePercent is not correct, value is %f", result.UsagePercent)
	}
	if result.RemainingPercent != 30 {
		t.Fatalf("Test Linux: RemainingPercent is not correct, value is %f", result.RemainingPercent)
	}
}
