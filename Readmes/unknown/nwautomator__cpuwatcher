# Readme #
This is a simple tool to watch CPU utilization on a Cisco device. The gauges are rendered using the Google Charts API. I created this as a way to easily watch CPU utilization during device testing, but it could come in handy for a variety of uses. 

## Requirements ##
To run this you'll need a webserver with at least PHP5 installed and enabled.

## Setup ##
Simply copy both the HTML and PHP files to your preferred location under your web server's root. Be sure to set the `$community` and `$device` variables in the PHP file as needed. Please note that the configured device is polled every quarter of a second. You should change the polling interval to fit your needs.

## Usage ##
With both files copied to your webserver and the appropriate values filled in for the two variables mentioned above, browse to `http://<your webserver>/cpuwatcher.html`. You should see three gauges which display the CPU utilization for the previous 5 seconds, 1 minute, and 5 minutes.
