package main

import (
	"ephemeral-prometheus-metric/commands"
	"ephemeral-prometheus-metric/logger"
	"ephemeral-prometheus-metric/metrics"
	"ephemeral-prometheus-metric/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	// Listen for keyboard interrupts
	utils.ListenForKeyboardInterrupted(nil)
	log := logger.GetLogger()
	port := utils.GetPort()
	diskPath := "/" // TODO: make it configurable

	promClient := prometheus.NewRegistry()
	m := metrics.NewMetrics(promClient, diskPath)
	result := commands.GetStorageData(diskPath)
	m.RootStorageTotalCapacityBytes.Set(result.TotalCapacityBytes)
	go func() {
		for {
			m.RootStorageRemainingBytes.Set(result.RemainingBytes)
			m.RootStorageRemainingPercent.Set(result.RemainingPercent)
			m.RootStorageUsagePercent.Set(result.UsagePercent)
			m.RootStorageUsageBytes.Set(result.UsageBytes)
			// update the data every 20 seconds
			time.Sleep(20 * time.Second)
			result = commands.GetStorageData(diskPath)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(promClient, promhttp.HandlerOpts{Registry: promClient}))
	log.Printf("Starting the HTTP server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
