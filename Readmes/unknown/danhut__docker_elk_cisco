Docker Compose configuration to bring up an ELK stack in support of Cisco equipment and Netrounds Virtual Test Agent.

Of note:
* containers are immutable - any config data should be baked into the container via the Dockerfile.  Run a  "docker-compose build" to re-bake the container if anything is changed

Logstash
* log_date is removed from the router log and replaces @timestamp to enable search on the device time

