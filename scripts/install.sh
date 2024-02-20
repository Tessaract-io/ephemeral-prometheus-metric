# Install the application by modifying the supervisor configuration file
echo "Installing the application"
mv ephemeral_prometheus_metric /usr/local/bin

# Modify the supervisord.conf file to include the application
echo "Modifying the supervisord.conf file"
echo -e "\n\n[program:ephemeral_prometheus_metric]" >> ../supervisord.conf
echo "command=/usr/local/bin/ephemeral_prometheus_metric" >> ../supervisord.conf

# Reload the supervisor configuration
echo "Reloading the supervisor configuration"
supervisorctl -c "../supervisord.conf" reread
supervisorctl -c "../supervisord.conf" update

# Start the application
echo "Starting the application"
supervisorctl start ephemeral_prometheus_metric
