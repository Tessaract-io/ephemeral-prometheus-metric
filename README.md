# Ephemeral Prometheus Metric
A simple script to expose a port and serve a Prometheus metric for any disk path usage.
The metric is a gauge and is updated every 20 seconds.

## Summary
> This code is written for the Linux OS, and Darwin only to scrape the `df -k` commands output then get the `Size`, `Used`, `Avail`, `Use%`, `Mounted on` columns for any partition.

## Usage
```bash
# Clone the repository
git clone <repo-url>

# Change directory
cd ephemeral-prometheus-metric

# Test and build the binary
./scripts/test_and_build.sh

# Run the binary directly, if <port> is not provided, it defaults to 15000
./ephemeral_prometheus_metric <port>

# Or update via supervisor:
#  A) Update the supervisor configuration with the script
./scripts/install.sh
#  B) Start the service
supervisorctl start ephemeral-prometheus-metric

# Check the result by visiting the following URL
curl http://localhost:<port>/
```
