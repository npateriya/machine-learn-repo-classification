acidoc
======

ACIDoc will document pertitent information about the configuration of a Cisco ACI fabric into Spreadsheet format,
it turns out sometimes that's still the best way to pass information around.

The contents and structure of the outputed format is very easy to customize
of the output is easily changed by modifying the [config.yml](config.yaml) file

## Prerequisites

Required

* Python 2.7+
* xlsxwriter
* pyyaml


You can install these requirements quickly by using

    pip install -r requirements.txt


 ## Configuration

 At a minimum you need to change the hostname and credentials in the config.yaml file to point to your
 APIC controller

```
host: http://apic1/
name: admin
passwd: password

```

You'll also likely want to change the filename which will be output by the tools

    filename: test.xlsx

## Running

Generating your spreadsheet is now a matter of running

    python aci-doc.py

