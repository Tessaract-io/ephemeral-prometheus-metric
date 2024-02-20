# Ephemeral Prometheus Metric
A simple script to expose a port and serve a Prometheus metric for any disk path usage.
The metric is a gauge and is updated every 20 seconds.

## Summary
> This code is written for the Linux OS, and Darwin OS only to scrape the `df -k` commands output then get the `Size`, `Used`, `Avail`, `Use%`, `Mounted on` columns for any wanted partition.

## Usage
```bash
# Clone the repository
git clone git@github.com:Tessaract-io/ephemeral-prometheus-metric.git

# Change directory
cd ephemeral-prometheus-metric

# Test and build the binary
/bin/ash ./scripts/test_and_build.sh || /bin/bash ./scripts/test_and_build.sh

# Run the binary directly, if <port> is not provided, it defaults to 15000
./ephemeral_prometheus_metric <port>

# Or update via supervisor
# Update the supervisor configuration with the script then start the service
/bin/ash ./scripts/install.sh || /bin/bash ./scripts/install.sh

# Check the result by visiting the following URL
curl http://localhost:<port>/
```
## Example Output
```bash
# HELP pod_storage_root_remaining_bytes Remaining Capacity of the / (root) storage.
# TYPE pod_storage_root_remaining_bytes gauge
pod_storage_root_remaining_bytes 1.57832237056e+11
# HELP pod_storage_root_remaining_percent Remaining Capacity of the / (root) storage in percentage.
# TYPE pod_storage_root_remaining_percent gauge
pod_storage_root_remaining_percent 58.58060132481901
# HELP pod_storage_root_total_capacity_bytes Total Capacity of the / (root) storage.
# TYPE pod_storage_root_total_capacity_bytes gauge
pod_storage_root_total_capacity_bytes 2.69427478528e+11
# HELP pod_storage_root_usage_bytes Usage of the / (root) storage in bytes.
# TYPE pod_storage_root_usage_bytes gauge
pod_storage_root_usage_bytes 1.11595241472e+11
# HELP pod_storage_root_usage_percent Usage of the / (root) storage in percentage.
# TYPE pod_storage_root_usage_percent gauge
pod_storage_root_usage_percent 41.41939867518099
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0
```