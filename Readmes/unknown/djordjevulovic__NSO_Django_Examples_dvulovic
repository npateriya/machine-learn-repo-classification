# Demo Northbound Applications for Cisco Network Services Orchestrator (NSO) built with Django Framework

## Introduction

[Cisco Network Services Orchestrator (NSO) enabled by Tail-f®] (http://www.cisco.com/c/en/us/solutions/service-provider/solutions-cloud-providers/network-services-orchestrator-solutions.html) is an industry-leading orchestration platform for hybrid networks. One of the main advantaes of NSO is comprehensive set of northbound APIs. With those APIs, NSO can be easily integrated with northbound systems such as operation support systems (OSSs) and self-service portals.

In this project, we will provide demos of northbound applications integrated with Cisco NSO using REST API. Applications are created by [Django framework] (https://www.djangoproject.com/), providing easy-to-use Web interface. 

List of applications:
- "NSO_Intro_Webdemo" provides introductory app which will show yoy how to use basics of NSO REST API. See app section below for details.

## Instructions (for all applications)

To run an application, first run through the following steps (skip the ones which you have already performed):
- Applications are designed to use NSO instance from [Cisco dCloud] (https://dcloud.cisco.com/) lab ["Cisco NSO Lab - L2VPN Service Design with XML Templates v1"] (https://dcloud2-lon.cisco.com/content/demo/67315). Make sure you do the following (all these steps are thoroughly described in the dCloud lab guide):
  - Register for Cisco.com account (http://www.cisco.com/web/siteassets/account/index.html)
  - Schedule lab session
  - Connect to VPN using Anyconnect
- Install Python 3 and Django (https://docs.djangoproject.com/en/1.10/intro/install/)
- Install Git client (https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- Clone project from the Github (`git clone https://github.com/djordjevulovic/NSO_Django_Examples_dvulovic`)
- Run Django server on port 8888 (`python manage.py runserver 8888`)
- Go to your preferred browser (tested with Chrome) and enter URL of the application main page (e,g. "http://127.0.0.1:8888/nso_intro/" for "NSO_Intro_Webdemo" application)

## Instructions (for "NSO_Intro_Webdemo" application)

Start this app by puting URL "http://127.0.0.1:8888/nso_intro/" into your favorite browser (tested with Chrome).

This simple app has three parts (pages):
- "Show Devices Configuration" will show full configuration of a device in JSON format. Input form contains combo-box for choosing device, which will be pre-propulated with evices available to NSO.
- "Show Interface Configuration" will configuration of a device interface in JSON format. Input form contains combo-boxes for available device and then available interface types and interface numbers for this specific device. When changing device, interface type/number boxes will automatically change.
- "Add New DNS Service" will deploy instance of a simple service for putting new DNS server configuration. Input form contains text box for service name, combo box for device, and text box for IP address of the DNS server.

In order to deploy a service, it must be defined and listed as a NSO package. Follow these steps to deploy DNS service (prerequsite for the "Add New DNS Service" part):
- Login to NSO server by doing SSH to 198.18.1.79 port 22 (username:cisco, password: C1sco12345)
- Change directory to "ncs-run/example1/packages"(`cd ncs-run/example1/packages`)
- Create skeleton of the new service package "DNS" (`ncs-make-package --service-skeleton template-based DNS`)
- Replace file "ncs-run/example1/packages/DNS/src/yang/DNS.yang" with file [DNS.yang] (https://github.com/djordjevulovic/NSO_Django_Examples_dvulovic/blob/master/NSO_Files/DNS%20Service/DNS.yang); use text editors like "vi" or "nano" or SCP/SFTP client.
- Replace file "ncs-run/example1/packages/DNS/templates/DNS.xml" with file [DNS.xml] (https://github.com/djordjevulovic/NSO_Django_Examples_dvulovic/blob/master/NSO_Files/DNS%20Service/DNS.xml); use text editors like "vi" or "nano" or SCP/SFTP client.
- Go to the "src" directory (`cd ncs-run/example1/packages/DNS/src`)
- Compile the service (`make`)
- Login to NSO by doing SSH to 198.18.1.79 port 2024 (username:admin, password: admin)
- Reload packages (`packages reload`)
- Check if DNS service is operational (`show packages package DNS`)

## Developer notes

- All wrapper functions for NSO REST API are in the [nso_helper_dvulovic.py] (https://github.com/djordjevulovic/NSO_Django_Examples_dvulovic/blob/master/nso_helper_dvulovic.py) file.
- Project is built using the PyCharm IDE
