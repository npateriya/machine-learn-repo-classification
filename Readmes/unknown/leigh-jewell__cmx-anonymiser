# cmx-anonymiser
This program uses an API to pull wireless client and Access Point (AP) data
from a Cisco CMX appliance. The Cisco CMX polls Cisco Wireless LAN Controllers
for metrics on wireless clients that are heard by wireless Access Points (APs).
Information such as RSSI and bytes sent/received can be interesting to
use for wireless performance metrics on your network.

As some of the data contains personal information such as username, ip address
and mac-address the program anonymises any personal information using a one-way
SHA256 hash. This allows for the data to be tokenised but anonymous to
ensure privacy is maintained. The results are stored in csv files for later
processing.
## Getting started
Simply create a directory and clone the git respository:   

> git clone https://github.com/leigh-jewell/cmx-anonymiser

This will download two files:   
> cmx-anonymiser.py - this is the python script that you will run   
> config.ini - this file contains all the configurable settings for the script  

When you run the program it will create two more directories:  
> output - where the csv files are written to  
> logs - all the logs from the running of the script  

## Prerequisites
The script has been written for Python 3 and tested on version 3.6.2. There are
a number of modules which you will need to have installed into your python
installation:

| Module        | Purpose                                         |
| ------------- |-------------------------------------------------|
|configparser   | Read in the config file                         |
|requests       | Nice URL module to access CMX API               |
|requests.auth  | The HTTP authentication part of requests        |
| collections   | Default collections                             |
| csv           | To write out the csv files                      |
| hashlib       | Hasing the mac address                          |
| datetime      | To calculate when to schedule the jobs          |
| os            | To work out the directories to use for output   |
| sched         | For scheduling the jobs to run                  |
| time          | For time                                        |
| math          | For math calculations                           |

You can simply install them with:
> pip install *module name*
## Installation and setup
The config.ini file contains all the configurable settings to control how
this script will run.
### CMX Hostname, username and password
The very minimum changes you will need to make will be to point the script
to your CMX server either by ip address or hostname and add in
a username/password to access it. You will need create an account on your
CMX so the script can connect via the API. By default the Cisco DevNet
sandbox CMX and username/password has been added so you can test it out.
### Output Directory
By default the output will be written to the current working directory output
and logs folder. You can change this behaviour to write it somewhere else. You
can also set log_console to tell the script to write logs to the console
if you are testing it out.
### API URLs
There are two API's which are used and this can be changed to something else.
This is more for when the CMX code is changed and you need to point it to a new
API in the future. The code is specifically written to look for known JSON
responses for these API calls to get the current active clients and aps. If you
use another different API you will need to modify the code to work. Possibly
in the future would be to write code that automatically puts the API return
into a CSV file.
### Schedule
You tell the program how many days and how often you want to poll the CMX. The
days is simply the number of days to poll and hours is a list of 24hr times to
get the data from CMX. Keep in mind the API calls will impact the CMX so
don't add to many in.
### Privacy and anonymisation
Most of the personal information is discarded that is returned from the CMX but
we need to retain the mac-address of the client in some form so we can
connect all the returned data to the same client. To avoid any personal
information from been collected the mac-address is converted to a one-way hash
using SHA256. To avoid collisions and ensure the mac-address remains private
a salt is also added to the string. This ensures that if you give out
the CSV files to a 3rd party that they can't easily brute force the mac addfress.
Obviously keep the salt private when sharing the CSV file. The salt
is just a string so make it whatever you want.

| Config           | Purpose                                            |
|------------------|----------------------------------------------------|
| cmx_ip           |Change it to your CMX MSE IP address or hostname    |
| username         | Username that exists on the CMX                    |
| password         | Password for the account on the CMX                |
| timeout          | How long to wait until timeout to CMX              |
| retry            | How many times to retry on connection failure      |
| retry_sleep      | How long to sleep before trying again after failure                                   |
| output_dir       | Directory to write the csv files, default output                                      |
| log_dir          | Log file directory, default logs                                                      |
| log_console      | Log to console, default True                                                          |
| url_clients      | Client API URL, default: /api/location/v2/clients                                     |
| url_client_count | Client API URL, default: /api/location/v2/clients/count                               |
| page_size        | How many clients to get each request, max 1,000  supported on CMX                     |
| max_pages        | Maximum client pages to pull back, just in case CMX client count return is very large |
| url_aps          | AP API URL, default: /api/config/v1/aps                                               |
| days             | How many days to collect, default 7                                                   |
| hours            | 24hr times to run default: 9:00,12:00,15:00,18:00                                     |
| hours            | Optional to set hours to 'now' to run the script right now                            |
| salt             | Random string to avoid hash collisions                                                |

## Running a test
Even without changing the config.ini you can test out the code as the default config
file is pointing to a Cisco CMX server in a DevNet sandbox and as long as this
server is reachable the test should run.
> python cmx-anonymiser.py

If console logging is enabled you will see the output on the console as it
pulls down the data and writes the csv files.

Example output:
> $ python cmx-anonymiser.py   
> 29/08/17 10:57.44.138577: main: Process started, scheduling jobs 7 days and now hours   
> 29/08/17 10:57.44.166156: getData: Process woken up.   
> 29/08/17 10:57.44.167660: getData:Using CMX:cmxlocationsandbox.cisco.com and username:learning   
> 29/08/17 10:57.44.169164: getCMXAPData: Getting data for: http://cmxlocationsandbox.cisco.com/api/config/v1/aps   
> 29/08/17 10:57.46.298581: getCMXAPData: Got status code 200 from CMX API (200 is good)   
> 29/08/17 10:57.46.298581: getCMXAPData: Got 9 ap records from CMX.   
> 29/08/17 10:57.46.298581: writeFile:Using output as output directory   
> 29/08/17 10:57.46.333839: writeFile:Finished writing.   
> 29/08/17 10:57.46.336848: getCMXData: Getting data for:http://cmxlocationsandbox.cisco.com/api/location/v1/clients   
> 29/08/17 10:57.51.102041: getCMXData: Got status code 200 from CMX API (200 is good)   
> 29/08/17 10:57.51.120462: getCMXData: Got 79 records from CMX.   
> 29/08/17 10:57.51.120462: writeFile:Using output as output directory   
> 29/08/17 10:57.51.149306: writeFile:Finished writing.   
> 29/08/17 10:57.51.151312: getData: Process sleeping.   
> 29/08/17 10:57.51.152316: main: Finished scheduled runs.   

### Output CSV files
Refer to the CMX API for details on what each field represents:
[CMX API 10.3 ref](https://www.cisco.com/c/en/us/td/docs/wireless/mse/10-3/api/b_cmx_103_api_reference/location.html)

#### File: user_data
This will be in CSV format but I have written it here by rows for easy reading.
> hash = abfc33006cfc08577dbc697540c0dbcd2cd8c962699b422e9d7cb677e537490f9   
> mapHierarchyString = DevNet  
> floorRefId = 723413321239068650   
> length = 74.1     
> width = 39.0   
> x = 12.29309   
> y = 67.07972   
> unit = FEET   
> currentlyTracked = True   
> confidenceFactor = 64.0   
> currentServerTime = 2017-08-29T01:57:48.292+0100   
> firstLocatedTime = 2017-08-16T16:05:05.953+0100   
> lastLocatedTime = 2017-08-29T01:57:48.115+0100   
> maxDetectedRssiApMacAddress = 0f:3b:01:c0:08:20   
> band = IEEE_802_11_B   
> rssi = -52   
> lastHeardInSeconds = 0   
> networkStatus = ACTIVE   
> changedOn = 1503968268115   
> ssId = test   
> band = IEEE_802_11_B    
> apMacAddress = 0f:3b:01:c0:08:20   
> dot11Status = ASSOCIATED   
> manufacturer = Trw   
> detectingControllers = 1.1.1.1
> bytesSent = 110   
> bytesReceived = 100   

#### File: ap_data
> radioMacAddress = 0f:3b:01:c0:08:20   
> name = T1-7   
> x = 3.5704226   
> y = 69.43445   
> unit = FEET   
> 802_11_BChannelNumber = 1   
> 802_11_BTxPowerLevel = 1   
> 802_11_AChannelNumber = 64   
> 802_11_ATxPowerLevel = 5   
> floorId = 723413320329068650   
### Issues
Testing on CMX 10.3, I was not consistently able to get results from the
API call /api/location/v2/clients. Sometimes I would get a 204 result
returned. I changed the API to /api/location/v1/clients and now
consistently get data returned. I can't find any documentation
to explain this but it works at the moment.
### Authors
Leigh Jewell
### License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/leigh-jewell/cmx-anonymiser/blob/master/LICENSE)
file for details.
### Acknowledgments
Thanks to [Cisco DevNet sample code cmx ](https://github.com/CiscoDevNet/sample-code-cmx)
for their CMX sample code repository.
