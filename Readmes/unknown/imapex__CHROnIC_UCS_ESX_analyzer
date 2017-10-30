# CHROnIC UCS ESX Analyzer
![logo](https://github.com/imapex/CHROnIC/blob/master/images/CHROnIC_logo_med.png)

Cisco Health Reports and Online Information Center is a plug-able web app to help visualize and maintain device's support by providing various health reports.

CHROnIC UCS ESX Analyzer uses information collected from CHROnIC Collector to determine if a UCS system is in compliance with Cisco's HCL

## Example Report
![](images/example_report.png)

# Installation

CHROnIC UCS ESX Analyzer is built in Python 3.5 and is packaged with a Dockerfile for easy builds and deployments

## Environment

CHROnIC UCS ESX Analyzer expects the following environment variables to be set:

CHRONICBUS needs to point to the CHROnIC Bus instance's base URL
```
CHRONICBUS
```

CHRONICPORTAL needs to point to the CHROnIC Portal instance's base URL
```
CHRONICPORTAL
```

HCL needs to point to one of Cisco's HCL tool URL's.  ex: http://ucshcltool.cloudapps.cisco.com/public/rest
```
HCL
```



# Usage
## python
```
pip install -r requirements.txt
export CHRONICBUS=<bus url>
export CHRONICPORTAL=<portal url>
export HCL=<hcl url>
python3 main.py
```

or

## Docker
```
docker build -t chronic_ucs_esx_analyzer
docker run -d -p 5000:5000 -e "CHRONICBUS=<bus url>" -e "CHRONICPORTAL=<portal url>" -e "HCL=<hcl url>" chronic_ucs_esx_analyzer
```

# Compatibility

Currently CHROnIC UCS ESX Analyzer needs some manual translations to be setup to work between UCS and the Cisco HCL tool.  These bindings live in a TinyDB file of piddb.json .  Items can be added at anytime by editing that database

Supported Servers:
* UCS B200 M3

Supported CPU's:
* E5-2600 v1
* E5-2600 v2

Supported Adapters:
* VIC 1240

Supported OS:
* VMware ESX (unknown if update versions run correctly)

# Suggested Next steps
* Change from TinyDB to a simple YAML file
* Add more pid to HCL descriptions
* Add better error handeling
