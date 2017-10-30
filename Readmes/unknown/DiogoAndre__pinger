# Pinger

## What

Simple **expect** script that accepts a list of hosts from a csv file, access a network device (Cisco), ping each host from the device and generates a semi-raw log file.

## Why

I needed to ping various locations from once from one site where we didn't have any servers that could do the job.

## How

The script takes two arguments. The first one is the IP address of the device it should access, and the second is the csv file with the list of hosts it should ping.

`./pinger.sh 10.1.1.1 hosts.csv`

The script will prompt for the username and password and password it should use.

After the script finishes you can run `parser.rb` to parse the log file and generate a csv.

**hosts.csv** is an example of csv file the script expects.

**20141027133023.log** is an example of the log file the script generates.

**csv-20141027133023.log.csv** is an example of the resulting csv file after running `parser.rb`.