# C H I V E   Agent Service
Cisco Heat Indication Visualization Explorer Agent

# Background
High temperatures are the enemy of network equipment. The traditional approach of setting thresholds and relying on SNMP traps to alert on environmental changes is very reactive. 
This application leverages the Application Policy Infrastructure Controller (APIC) REST APIs to continuously gather near real-time operating temperatures of both spine and leaf switches in the Cisco Application Centric Infrastructure (ACI) fabric. 
The data that is collected is then displayed on a webpage. The eventual goal of the project is to create a data center "heatmap" to show hotspots on the data center floor.

Devices are categorized into three conditions HIGH, ELEVATED and NORMAL. Each condition is defined by a temperature range and is assigned a color

(1) HIGH     -  red   - the device has exceeded safe operating temperatures 
(2) ELEVATED - yellow - device temperature had gone above normal temperature range
(3) NORMAL   - green  - device is operating within expected temperature range

# Application Details
This application was developed based on microservice architecture and is wrapped in a Docker container.

One image is designed so that it can be deployed on a ARM (Raspberry Pi) device. Another image exists that can be deployed on an x86 device. The user is able to choose the image at installation time. 

The application invokes an API GET request by sending an HTTP/1.1 GET message to the Application Policy Infrastructure Controller (APIC). The HTML body of the response message from the APIC controller contains a Javascript Object Notation (JSON) structure that contains the requested temperature data.
Once the data that was received is parsed, the JSON structured data is sent to the **chive_app** microservice ,running in the MANTL environment, via an API POST request. 

The HTML body of the POST message is JSON.

# Diagram

Inline-style: 
![alt text]( https://github.com/imapex/chive_agent/blob/master/diagrams/CHIVE_AGENT_REST_calls.gif "CHIVE_AGENT_REST_calls")

