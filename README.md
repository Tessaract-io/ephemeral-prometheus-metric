# Ephemeral Prometheus Metric
A simple script to expose a port and serve a Prometheus metric for any disk path usage.
The metric is a gauge and is updated every 20 seconds.

## Summary
> This code is written for the Linux OS, and Darwin only to scrape the `df -k` commands output then get the `Size`, `Used`, `Avail`, `Use%`, `Mounted on` columns for any wanted partition.

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
