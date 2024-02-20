package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type Metrics struct {
	RootStorageTotalCapacityBytes prometheus.Gauge
	RootStorageRemainingBytes     prometheus.Gauge
	RootStorageRemainingPercent   prometheus.Gauge
	RootStorageUsagePercent       prometheus.Gauge
	RootStorageUsageBytes         prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer, diskPath string) *Metrics {
	description := fmt.Sprintf("%s storage", diskPath)
	diskName := strings.Replace(diskPath, "/", "-", -1)
	if diskPath == "/" {
		description = "/ (root) storage"
		diskName = "root"
	}
	m := &Metrics{
		RootStorageTotalCapacityBytes: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("pod_storage_%s_total_capacity_bytes", diskName),
			Help: fmt.Sprintf("Total Capacity of the %s.", description),
		}),
		RootStorageRemainingBytes: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("pod_storage_%s_remaining_bytes", diskName),
			Help: fmt.Sprintf("Remaining Capacity of the %s.", description),
		}),
		RootStorageRemainingPercent: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("pod_storage_%s_remaining_percent", diskName),
			Help: fmt.Sprintf("Remaining Capacity of the %s in percentage.", description),
		}),
		RootStorageUsagePercent: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("pod_storage_%s_usage_percent", diskName),
			Help: fmt.Sprintf("Usage of the %s in percentage.", description),
		}),
		RootStorageUsageBytes: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("pod_storage_%s_usage_bytes", diskName),
			Help: fmt.Sprintf("Usage of the %s in bytes.", description),
		}),
	}
	reg.MustRegister(m.RootStorageTotalCapacityBytes)
	reg.MustRegister(m.RootStorageRemainingBytes)
	reg.MustRegister(m.RootStorageRemainingPercent)
	reg.MustRegister(m.RootStorageUsagePercent)
	reg.MustRegister(m.RootStorageUsageBytes)
	return m
}
