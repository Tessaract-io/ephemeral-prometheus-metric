# Install the application by modifying the supervisor configuration file
echo "Installing the application"
if [ -z "$1" ]; then
  pass
  elif [ "$1" = "alpine" ]; then
    echo "Copying the application to /usr/local/bin/ephemeral_prometheus_metric for alpine"
    cp ./bin/alpine/ephemeral_prometheus_metric /usr/local/bin/ephemeral_prometheus_metric
else
  echo "Copying the application to /usr/local/bin/ephemeral_prometheus_metric for ubuntu"
  cp ./bin/ubuntu/ephemeral_prometheus_metric /usr/local/bin/ephemeral_prometheus_metric
fi

# Modify the supervisord.conf file to include the application
program_exists=$(grep -c "ephemeral_prometheus_metric" ../supervisord.conf)
if [ $program_exists -eq 0 ]; then
    echo "Modifying the supervisord.conf file"
    echo -e "\n\n[program:ephemeral_prometheus_metric]" >> ../supervisord.conf
    echo "command=/usr/local/bin/ephemeral_prometheus_metric" >> ../supervisord.conf
else
    echo "The supervisord.conf file is updated already"
    exit 0
fi

# Reload the supervisor configuration
echo "Reloading the supervisor configuration"
supervisorctl -c "../supervisord.conf" reread
supervisorctl -c "../supervisord.conf" update
