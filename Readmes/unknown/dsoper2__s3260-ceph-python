# s3260-ceph-python
Python based deployment of Cisco UCS S3260 Storage Server with Red Hat Ceph Storage

#### Table of Contents

1. [Description](#description)
1. [Installation - The basics of getting started with s3260-ceph-python](#installation)
    * [What s3260-ceph-python affects](#what-s3260-ceph-python-affects)
    * [Beginning with s3260-ceph-python](#beginning-with-s3260-ceph-python)
1. [Usage - Configuration options and additional functionality](#usage)

## Description

Provides a reference deployment of portions of the UCS configuration outlined in the "Cisco UCS S3260 Storage Server with Red Hat Ceph Storage" Cisco Validated Design available here: http://www.cisco.com/c/en/us/td/docs/unified_computing/ucs/UCS_CVDs/ucs_s3260_rhceph21.html

## Installation

1.  Clone the s3260-ceph-python module into a working directory of your choice:
    ```
    git clone https://github.com/dsoper2/s3260-ceph-python.git
    ```

### What s3260-ceph-python affects

* Dependencies that s3260-ceph-python requires.

s3260-ceph-python depends on the UCS Central Python SDK.  If the UCS Central Python SDK is not installed, please see https://github.com/CiscoUcs/ucscsdk for installation instructions.

## Usage

* Python scripts included in s3260-ceph-python general have the following usage:
    ```
    python ucsc_<action> <JSON credentials file> <csv file with parameters>
    ```

* The JSON credentials file should be customized for your UCS Central enviornment.  An example file, ucsc_dcloud.json, is included with the following content that specifies IP address, username, and password:
    ```
    {
    "ip": "ucs-ctl1.dcloud.cisco.com",
    "pw": "C1sco12345",
    "user": "admin"
    }
    ```

The csv file provides parameters for Service Profile Template, org, name, and other pools/policies/profiles used by the Template.  The csv file can also include Service Profile names and number of instances used in creating Service Profiles from the template.  Use of a csv file for parameters is provided only as an example and is not required by any parts of the Python SDK.  See global_spt_test.csv for example parameters.
* With the JSON credentials file and csv file customized for your environment, here are specific command lines to create and delete Global Service Profile Templates:
    ```
    python ucsc_create_global_spt.py ucsc_dcloud_demo.json global_spt_test.csv 
    
    python ucsc_delete_global_spt.py ucsc_dcloud_demo.json global_spt_test.csv 
    ```

* Here are example command lines to associate and disassociate Global Service Profile Templates:
    ```
    python ucsc_associate_sp.py ucsc_dcloud_demo.json global_spt_test.csv 
    python ucsc_disassociate_sp.py ucsc_dcloud_demo.json global_spt_test.csv 
    ```
