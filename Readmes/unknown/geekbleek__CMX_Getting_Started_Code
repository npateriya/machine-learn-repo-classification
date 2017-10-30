# CMX Quick Start Example Code

This repository contains sample code to get started with the Cisco CMX Wireless Location Analytics service in just a few minutes.  It will allow you to:
  - Connect to a Cisco CMX Instance
  - Retrieve the location of a device via mac address
  - Get the details of that device's location including campus, building, and floor information
  - Retrieve an the image and dimensions of the floorplan where the device is located
  - Leverage native image libraries to generate a map showing where a device is currently located
  
The sample code is provided for Node.js & Python (version 3.6 & 2.7)

# Basic Usage

The following environment variables are expected to be set, either as system environment variables or at runtime:
  - macAddress - the macAddress of the device you want to locate such as '00:00:2a:01:00:27'.  For more info see "Finding macAddress" below)
  - baseCMXUrl - the URL of the CMX instance i.e. 'https://cmxlocationsandbox.cisco.com'
  - CMXusername - the CMX username
  - CMXpassword - the CMX password

Once executed, the application will output the map with the user's location into the /images directory.

# Node.js Usage

In the [node directory](/node) run

```
npm install
```

Then you can execute the sample.js file.  If your environment variables have already been set run:

```
node sample.js
```
Or to set them at runtime:
```
macAddress=00:00:2a:01:00:27 baseCMXUrl=https://cmxlocationsandbox.cisco.com CMXusername=learning CMXpassword=learning node sample.js
```
If successful you should see output similar to this:

```
Getting location data for 00:00:2a:01:00:27
Getting map for DevNetCampus>DevNetBuilding>DevNetZone>Zone1
Compositing map of location for (1372.4262424755702,156.24944683760683)
Location composited on map successfully.  Image output to ./images/composite.jpg
```

Note: This example uses Node.js promises - if you are using an older version of Node.js (prior to v4.0) that doesn't support ECMA you may want to consider using Bluebird for Promise support.
```
npm install bluebird
```
and then add to the sample.js:
``` js
var Promise = require("bluebird");
```

# Python Usage

Make sure the following Python packages are installed:

``` python
#Pillow is an open source native image manipulation library (https://github.com/python-pillow/Pillow)
pip install Pillow 

#Requests library
pip install requests
```

Then you can execute the sample.py file.  If your environment variables have already been set run:
```
python sample.py
```
Or to set them at runtime:
```
macAddress=00:00:2a:01:00:27 baseCMXUrl=https://cmxlocationsandbox.cisco.com CMXusername=learning CMXpassword=learning python sample.py
```

If successful you should see output similar to this:
```
Getting location data for 00:00:2a:01:00:27
Getting map for DevNetCampus>DevNetBuilding>DevNetZone>Zone3
Compositing map of location for (822.4250087296417,110.85201582417581)
Location composited on map successfully.  Image output to ./images/composite.jpg
```

Note: This example uses Python 3, but all of the dependencies are supported on 2.7 as well.  If using pyenv, the .python-version files are already included in this repo.  To run the python 2.7 version (the only change is print statement format) change to [python2.7 directory](./python/python2.7) and run the above commands.


# Finding macAddress of a device
If you are wanting to track your own device on the CMX infrastructure, you'll need to know your mac address! To find your device's mac address, use [this handy guide](https://kb.netgear.com/1005/How-to-find-a-MAC-address).


# Further CMX Resources
CMX has an API to list of all the devices available for tracking. For example:
```
curl -X GET \
  https://cmxlocationsandbox.cisco.com/api/location/v2/clients \
  -H 'authorization: Basic bGVhcm5pbmc6bGVhcm5pbmc='
```

You can also filter by UTC timestamps in MS to find historical data
```
curl -X GET \
  'https://cmxlocationsandbox.cisco.com/api/location/v1/history/clients?locatedAfterTime=1497225600000&locatedBeforeTime=1497376036000' \
  -H 'authorization: Basic bGVhcm5pbmc6bGVhcm5pbmc='
```

In addition, CMX allows you to register for Webhook notifications of events such as clients entering an area, movement events and much more!
[These CMX API's are all documented here.](https://cmxlocationsandbox.cisco.com/apidocs/)  There are also [4 DevNet Learning Labs](https://learninglabs.cisco.com/modules/dna-cmx-mse) to help you quickly learn all the features available on the CMX!

This sample code combined with the DevNet Learning Labs should allow you to quickly leverage powerful analytics from the Cisco CMX Wireless infrastructure!
