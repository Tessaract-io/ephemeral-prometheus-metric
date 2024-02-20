# Run the tests
echo "Running tests"
go test -v ./tests

# Build the application
echo -e "\n\nBuilding the application"
go build -o ./ephemeral_prometheus_metric main.go
echo "Application built, you can run it with ./ephemeral_prometheus_metric"
