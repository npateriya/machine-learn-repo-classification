# CHIVE Demonstration

Cisco Heat Indication Visualization Explorer is a sample application that illustrates the power and flexibility of the API and sensor data within a Cisco ACI fabric.  

## Demo Application Background

Cisco has provided with powerful APIs within many of the products they manufacturer, but most network engineers feel that programming and using those APIs is too difficult.  We set out to attack this by showing in a simple 3 tier application that we can use information within Cisco manufactured devices to provide us with a quick visualization of data.  In this case we are using the temperature sensors located within the ACI fabric.  The platform has the following capabilities:

* Capture temperature data from each switch node in an ACI Fabric.
* Provide a simple API Gateway to consume that information.
* Present that information onto a customizable web page that users are able "to position their "drag and drop" position the switches that have been collected from within their environment.

## Demo Application Requirements

* Docker must be installed
    * For installation instructions on installing Docker on Raspberry Pi click [here](http://blog.hypriot.com/getting-started-with-docker-on-your-arm-device/) and select the operating system you are on.
    * To install Docker for Mac or Linux [click here](https://www.docker.com/products/overview)
    * Currently Docker on Windows is not supported.

## Other Demo Repositories

This repository and README provide overall details on the Demonstration application and how to deploy the full demonstration.  

The following repositories are where the actual code for the different involved micro-services/components are located.  

* [chive_demo](https://github.com/imapex/chive_demo) - Full Demo Application Setup and Details
* [chive_app](https://github.com/imapex/chive_app) - Details on the CHIVE API Gateway 
* [chive_agent](https://github.com/imapex/chive_agent) - Details on the CHIVE Agent 
* [chive_web](https://github.com/imapex/chive_web) - Details on the User UI

# Demo Setup Steps

Within the CHIVE Demonstration there is a setup script that will deploy the three application components into a Docker environment.  The following pre-requisites must be met:

* Docker Environment must be setup on either a mac, linux, or RPi. 
	* No Windows support at this time.
* A working ACI fabric must be available with ethernet connectivity to the above mentioned Docker environment.
* User credentials for the ACI fabric must have API read privileges. 
	* Please note that version 1 of the CHIVE deployment uses http NOT https.  Therefore the APIC must be enabled to support http.

### Setup Instructions 

1. First use git clone the [chive_demo](https://github.com/imapex/chive_demo) repository on the device you are planning to install.
2. Type "cd chive_demo"
3. Run the "deploy.sh" script.

You will be prompted to enter the following:
1. APIC IP Address (FQDN is an option if DNS is available)
2. APIC Username with API access
3. APIC Password for the above username

## User Interaction 

Upon the successful deployment of the above application you should be able to run "docker ps" from the cli and see 3 running containers.

If all three containers are running then you will be able to web
http:<'IP Address of Deployment Device'>:5001  

The web page requires an initial setup of your environment such as dragging the respective APIC devices to their location on the provided grid layout.
	* Future versions will allow a custom image upload.


## Backlog

1. Custom Image Upload
2. Enable Https and more secure REST calls
3. Database component for data retention
4. Additional sensors
5. A more robust heatmap


