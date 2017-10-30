# ioxduino
IOx Proof of Concept to integrate a Cisco Router with an Arduino Microcontroller

![](resources/setup_image4.jpg)

### Table of Contents 

* [Background](#background)
* [Setup and PreRequisites](#setup-and-prerequisites)
* [Loading Demo Application](#loading-demo-application)
* [Running the Demo](#running-the-demo)
* [Known Caveats](#known-caveats)

## Background

Many IOT applications involve reading details from sensors, processing the information, and taking some action.  This action could be data storage, sending an alert notification, or updating some other system.  This a isn't new processing flow, what is new with IOT is that this interaction takes place connected to the Internet, or perhaps just an Intranet.  

A challenges for the IOT application developer are: 

1. How to bridge the gap from what has been the serial communciation world of sensors, and the IP based Internet.  
2. Where to run the applications that the monitor and interact with the external sensors.  

Cisco looks to solve these challenges with the [Cisco IOx](https://developer.cisco.com/site/iox/index.gsp) platform for Internet of Things applications.  Cisco has long been the leader in connecting people, applicaitons, business, etc to the Internet, and with IOx they now look to connect Things to the Internet.  

With IOx, application developers can run their IOT applications right at the edge of the network, close to the sensors and other devices (often called *Fog Computing*).  IOx Applications run directly on the Cisco Router that simulatneously provide the connectivity to the Internet or Intranet.  

## Arduino + Cisco IOx

[Arduino](https://www.arduino.cc) is an OpenSource, inexpensive, and easy to leverage microcontroller platform which can be easily used to interact with serial and digital components such as: 

* Temperature, Motion, Pressure, Light, etc Sensors 
* LED Lights for indictation
* Motors
* And many other "things"

Arduino is a great platform for building these electrical components and circuits, but lacks the horsepower to build complex applications.  

Cisco Routers and devices supporting IOx, provide the computing platform for running the application code, and includes serial (and other) interfaces for connecting to external devices like an Arduino.  

## ioxduino application 

ioxduino is provided as proof of concept code illustrating how such an application could be built leveraging IOx and an Arduino.  

In ioxduino, the external sensor running on the Arduino is a basic push button.  When the button is pressed the Arduino sends an event out the serial interface.  The serial interface on the Arduino is connected to a Serial Interface on the IOx Router, which is monitored by the IOx application code.  

When a button press is detected by the IOx application it updates a log that is available with a basic REST API call.  


# Setup and PreRequisites 

To build and replicate this demonstraiton you'll need access to an IOx device with a Serial Port, and a compatible Arduino device.  The original demonstration was created using a **Cisco C819HG-4G-G-K9** router and an **Arduino UNO**, however other models of each could be used.  

## IOx Router/Host 

[Cisco DevNet IOx Community](https://developer.cisco.com/site/iox/index.gsp) is the best place for resources on developing with IOx.  Before beginning this, or any other IOx app project, you can find information on preparing your router.  This includes: 

* installing the correct IOS image 
* configuing IOS to enable IOx 
* loading the neccessary cartridges for running applications

This demonstration runs as a Python PaaS application, so follow the [PaaS QuickStart](https://developer.cisco.com/media/iox-dev-guide-7-12-16/getstarted/quickstart-paas/) preperation for that demonstration.  

### Serial Port 

Because this demonstration uses the router's Serial Interface to connect to the Arduino, you'll need: 

* Additional configuration to enable the mapping of the Serial Interface to the IOx application
	* See this link for information [Cisco ISR 8xx - Enabling Serial Port](https://developer.cisco.com/media/iox-dev-guide-7-12-16/platforms/platform-isr8xx/#enabling-serial-port)
* A compatible Serial cable for your router that exposes the TX/RX pins to conenct to the Arduino.  

#### SmartSerial Interface Cable Option

The Cisco C819HG-4G-G-K9 router, and several others, leverage a SmartSerial interface on the Serial Interface.  To leverage this port, you can order a SmartSerial to RS232/DB25 cable that can be used to gain access to a standard serial interface.  That cable coupled with a DB25 breakout board makes it very easy to interconnect with the Ardunio.  

**Sample Parts List**

| Description | Part Number | Sample Link |
| ----------- | ----------- | ----------- |
| SmartSerial to Female DB25 | CAB-SS-232FC | [Amazon](https://www.amazon.com/Cisco-Smart-Serial-Female-CAB-SS-232FC/dp/B0064CQWDW/ref=sr_1_1?ie=UTF8&qid=1473424993&sr=8-1&keywords=smart+serial+to+db25) |
| Male DB25 Breakout Board | N/A | [Amazon](https://www.amazon.com/Swellder-Connector-25-pin-Terminal-Breakout/dp/B00V7S79BW/ref=sr_1_2?s=electronics&ie=UTF8&qid=1473425100&sr=1-2&keywords=db25+breakout) |

*Links to Amazon above are purely for reference.  We will receive no money if you use them.*

If you use a Female DB25 interface, the pins that you'll need are:

* **Pin 2 - Receive** connected to **RS232 Shifter Pin 2: Tx**
* **Pin 3 - Transmit** connected to **RS232 Shifter Pin 3: Rx** 
* **Pin 7 - Ground** connected to **RS232 Shifter Pin 5: Ground**

## Arduino 

The great part about the Arduino platform is the large number of boards to use that, for the most part, are compatible with each other, including sketchs (an Arduino program).  A common Ardunio to use for projects focused on learning and demonstrations is the Arduino UNO, which was used here.  

In addition to the Arduino board itself, you'll need a few other things to complete the external sensor circuit.  These include: 

* Push Button
* LED Light (for indication)
* Resistors (for LED and Pull Down on button)
* Breadboard (optional but very handy)
* Breadboard jumpers
* Arduino Power Supply

### Parts List 

The following parts list and links are provided for information and as a reference.  Any similar components will work.  

| Description | Link | 
| ----------- | ---- | 
| Arduino Uno | [Adafruit](https://www.adafruit.com/products/50) |
| Arduino Power Supply | [Adafruit](https://www.adafruit.com/products/63) |
| Arduino Uno USB Cable | [Adafruit](https://www.adafruit.com/products/62) |
| Breadboard | [Adafruit](https://www.adafruit.com/products/64) |
| Jumpers | [Adafruit](https://www.adafruit.com/products/153) |
| Button | [Adafruit](https://www.adafruit.com/products/367) |
| LED | [Adafruit](https://www.adafruit.com/products/299) |
| 2.2K Resistor - LED | [Adafruit](https://www.adafruit.com/products/2782) |
| 10K Resistor - Pull Down | [Adafruit](https://www.adafruit.com/products/2784) |

Alternatively, you can order an Arduino Starter Kit that includes the above and more.  

[Adafruit Arduino Starter Pack](https://www.adafruit.com/products/68)

*Links to Adafruit above are purely for reference.  We will receive no money if you use them.*

#### Arduino Serial TTL to RS232 

The Arduino Serial communciations operate using TTL, while the IOx router uses RS232.  These mechanisms are NOT directly compatible, and attempts to do so can damage your Arduino.  You can translate between them using a few different methods.  For this demo, I used an RS232 Shifter available from Sparkfun to make it simple and easy.  

| Description | Link | 
| ----------- | ---- | 
| RS232 Shifter | [Sparkfun](https://www.sparkfun.com/products/449) |
| DB9 Breakout | [Amazon](https://www.amazon.com/Swellder-Breakout-Terminals-Connector-Terminal/dp/B00Z2LIHAC/ref=sr_1_3?ie=UTF8&qid=1474033119&sr=8-3&keywords=db9+breakout) |

*Links to Sparkfun and Amazon above are purely for reference.  We will receive no money if you use them.*

### Component Diagram 

![](resources/arduino_diagram.png)


# Loading Demo Application

There are two steps to getting this demonstration up and working.  Loading and starting the ioxduino PaaS application on the IOx Device, and loading your Arduino with the sketch.  

Within this repository there are two folders with the needed code.  

* [ardunio_sketches](arduino_sketches/) contains the BasicButtonSerial code for the sample sensor event 
* [iox_app](iox_app/) contains the needed code to package and deploy ioxduino to your device.  

Begin by cloning the ioxduino repository to your local machine.  

```
git clone https://github.com/imapex/ioxduino 

# Change to the ioxduino directory
cd ioxduino/

# View the contents of the repo, you shoudl see something similar
ls
LICENSE			arduino_sketches       	resources
README.md      		iox_app

```

## IOx PaaS Application

Follow these steps to properly package and deploy ioxduino to your device.  This assumes you already have installed and setup ioxclient on your local machine.  For information on how to do that, see [ioxclient-reference](https://developer.cisco.com/media/iox-dev-guide-7-12-16/ioxclient/ioxclient-reference/) on DevNet.  

*All of these steps will be accomplished from a terminal on your workstation*

1. Enter the iox_app directory.  You should see the following files.  

	```
	cd iox_app/
	```
	
	| File | Description |
	| --- | --- | 
	| package.yaml | YAML formated description of the IOx application |
	| package_config.ini | Application configuration file (empty) | 
	| device_mapping.json | Host Router Device Mapping (ie network and serial) | 
	| requirements.txt | Python requirements list | 
	| main.py | Python code to run for application | 
	

2. Create an IOx application package that can be installed to your host device.  

	```
	ioxclient package .
	
	Currently active profile :  default
	Command Name: package
	Checking if package descriptor file is present..
	Created Staging directory at :  /var/folders/dh/t0frtgx514388l6ycj2bl7740000gn/T/238313704
	Copying contents to staging directory
	Checking for application runtime type
	Detected Python application. Attempting to install dependencies present in requirements.txt ..
	Collecting pyserial (from -r requirements.txt (line 1))
	  Using cached pyserial-3.1.1-py2.py3-none-any.whl
	Installing collected packages: pyserial
	Successfully installed pyserial
	
	Successfully installed dependencies..
	Creating an inner envelope for application artifacts
	Excluding  package.tar
	Generated  /var/folders/dh/t0frtgx514388l6ycj2bl7740000gn/T/238313704/artifacts.tar.gz
	Calculating SHA1 checksum for package contents..
	Root Directory :  /private/var/folders/dh/t0frtgx514388l6ycj2bl7740000gn/T/238313704
	Output file:  /var/folders/dh/t0frtgx514388l6ycj2bl7740000gn/T/290598695
	Path:  artifacts.tar.gz
	SHA1 : ebb9550385ea63c33565c3f8e18bed9dfe0dfbf8
	Path:  package.yaml
	SHA1 : 2ca92ad2d72465d42ebdbfd3c7db2a2aa01d0f3a
	Path:  package_config.ini
	SHA1 : da39a3ee5e6b4b0d3255bfef95601890afd80709
	Generated package manifest at  package.mf
	Generating IOx Package..
	Created IOx package at :  /tmp/ioxduino/iox_app/package.tar
	```

3. Install the ioxduino package onto your router. 

	```
	ioxclient application install ioxduino package.tar 
	
	Currently active profile :  default
	Command Name: application-install
	Installation Successful. App is available at : https://10.192.81.81:8443/iox/api/v2/hosting/apps/ioxduino
	Successfully deployed
	```
	
	* Verify application installed.
		
		```
		ioxclient application list 
		
		Currently active profile :  default
		Command Name: application-list
		List of installed App :
		 1. ioxduino   --->   DEPLOYED
		  
		```
	
4.  Update `device_mapping.json` Serial device name to match your router.  
	* Get the information about your router's interface. 

		```
		# List the devices avialable on your platform 
		ioxclient platform device list
		
		Currently active profile :  default
		Command Name: plt-device-list
		-------------Device Info----------------
		{
		 "serial": [
		  {
		   "available": true,				<--- Must show true
		   "device_id": "/dev/cpts0",		<--- This value for device_mapping.json
		   "device_name": "serial0",		<--- This indicates the router interface 
		   "port": 32000,
		   "slot": 0,
		   "used_by": null
		  }
		 ]
		}	
		```
	
	* Update `device_mapping.json` to use the correct serial_id.
	
		```
		cat device_mapping.json
		{
		  "resources": {
		    "cpu": "25",
		    "devices": [
		      {
		        "device-id": "/dev/cpts0",	<--- Update this to match
		        "label": "HOST_DEV1",
		        "type": "serial"
		      }
		    ],
		    "disk": "100",
		    "memory": "50",
		    "network": [
		      {
		        "interface-name": "eth0",
		        "network-name": "iox-bridge0"
		      }
		    ],
		    "profile": "custom"
		  }
		}	
		```

5. Activate ioxduio application 

	```
	ioxclient application activate ioxduino --payload device_mapping.json
	
	Currently active profile :  default
	Command Name: application-activate
	Payload file : device_mapping.json. Will pass it as application/json in request body..
	App ioxduino is Activated	
	```
	
	* Check that ioxduio is now activated
	
		```
		ioxclient application list
		Currently active profile :  default
		Command Name: application-list
		List of installed App :
		 1. ioxduino   --->  ACTIVATED
		```

6. Start ioxduino application.  

	```
	ioxclient application start ioxduino 
	
	Currently active profile :  default
	Command Name: application-start
	App ioxduino is Started
	```
	
7. Get the IP address assigned to ioxduino
		
	```
	# Run this command to get the details of the started application
	# If no ipv4 address is displayed, wait a minute and re-run
	ioxclient application info ioxduino 
	
	# Some output not displayed 
	Currently active profile :  default
	Command Name: application-info
	Details of App : ioxduino
	-----------------------------
	{
	 "description": "Monitors External Arduino Sensor and writes collected sensor data to log files",
	 "name": "ioxduino",
	 "networkInfo": {
	  "eth0": {
	   "ipv4": "192.168.1.16",				<--- This is the IP Address for ioxduino
	   "ipv6": null,
	   "libvirt_network": "dpbr_0",
	   "mac": "52:54:DD:EE:D7:BF",
	   "mac_address": "52:54:dd:ee:d7:bf",
	   "network_name": "iox-bridge0",
	   "port_mappings": {
	    "tcp": [
	     [6000,6000]							<--- The TCP port for ioxduino 
	    	]
	   }
	  }
	 },
	 "resources": {
	  "cpu": 25,
	  "devices": [
	   {
	    "device-id": "/dev/cpts0",		<--- Note the Serial Interface
	    "label": "HOST_DEV1",
	    "type": "serial",
	    "usage": "Integrating with Arduino"
	   }
	  ]
	 },
	 "state": "RUNNING",
	}	

	```

8. If you've configured your IOx Platform in the default mode where the IOx Guest OS is NAT'd to the WAN interface, you'll need to install a Static NAT entry for port 6000 to the ioxduino IP Address assigned.  If your IOx Platform networking is configured in another way, you can skip this part.  
	1.  Login to IOS running on your router and enter config mode.  

		```
		iox#conf t
		
		! Replace 192.168.1.16 with the IP address assigned to ioxduio 
		! If your Outside interface is NOT GigabitEthernet 0, replace it
		iox(config)#ip nat inside source static tcp 192.168.1.16 6000 interface GigabitEthernet0 6000
		iox(config)#exit
		
		! Verify the NAT Translation was configured 
		iox#show ip nat translations		
		Pro Inside global         Inside local          Outside local         Outside global
		tcp <- RTR IP ->:6000     192.168.1.16:6000     ---                   ---
		tcp <- RTR IP ->:2222     192.168.3.2:22        ---                   ---
		tcp <- RTR IP ->:8443     192.168.3.2:8443      ---                   ---		
		```

9. With ioxduino started, and NAT configured, you can now access the REST API to get a list of alerts.  

	```
	# Replace RTR_IP with your router's IP address
	curl RTR_IP:6000/
	
	[]
	
	# Currently no alerts have been logged
	```
	

## Arduino Sample Sketch 

Before loading the Sample Sketch onto your Arduino, be sure to have built the electrical circuit as shown in the [Component Diagram](#component-diagram).  

*See []() for information on full details using the Arduino IDE.  The steps below assume some familiarity with the process of programming and Arduino*

1. Open `BasicButtonSerial.ino` in the Arduino IDE.  
2. Make sure your Ardino board and port are configured.  
3. Upload the Sketch to your Arduino.  

	![](resources/arduino_upload.png)

4. Verify that Arduino is configured properly by opening the Serial Monitor within the IDE and pushing the button.  You should see a message for each button push.  

	![](resources/arduino_serial_monitor.png)

5. Disconnect the USB Cable from the Arduino used to program.  Be sure another power adapter or supply is connected still.  

# Running the Demo

With the both the PaaS application and the Arduino Sketch loaded and tested individually, now you will connect the two devices together for the full demo.  

## Connecting the Serial Devices Together

Depending on the exact type of Serial interface and cable you are using with your router, you may need to update these details.  But if your router serial cable provides a Female DB25 connector (as was used in this demo build) you will use these pin connections.  

* **Pin 2 - Receive** connected to **RS232 Shifter Pin 2: Tx**
* **Pin 3 - Transmit** connected to **RS232 Shifter Pin 3: Rx** 
* **Pin 7 - Ground** connected to **RS232 Shifter Pin 5: Ground**

## Testing the connection 

Once connected, you can press the button on the Arduino a few times (pause at least a second between presses) to generate some alerts.  Then make an REST API call to ioxduino to see if you're logging alerts.  

```
curl RTR_IP:6000/
[["Fri Sep 16 12:39:52 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:05 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:07 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:09 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:12 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:13 2016", "Button 6 pressed."], ["Fri Sep 16 12:40:17 2016", "Button 6 pressed."]]

# Pass the returned results through json.tool to pretty the display some 
curl 10.192.81.81:6000/ | python -m json.tool 

[
    [
        "Fri Sep 16 12:39:52 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:05 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:07 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:09 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:12 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:13 2016",
        "Button 6 pressed."
    ],
    [
        "Fri Sep 16 12:40:17 2016",
        "Button 6 pressed."
    ]
]
```

You can also use ioxclient to connect to the console of the running app to monitor the output from the router.  

```
ioxclient app console ioxduino

# Some output removed
Currently active profile :  default
Command Name: application-console
.
.
Connected to domain ioxduino
.
Poky (Yocto Project Reference Distro) 1.7.2 isr800-lxc /dev/tty1

isr800-lxc login: Serving on port 0.0.0.0:6000
Serial:  Serial<id=0x10144110, open=True>(port='/dev/cpts0', baudrate=9600, bytesize=8, parity='N', stopbits=1, timeout=5, xonxoff=False, rtscts=False, dsrdtr=False)

10.192.81.112 - - [11/Sep/2016 20:13:39] "GET / HTTP/1.1" 200 2
Incoming Data found.
['Fri Sep 16 12:39:52 2016', 'Button 6 pressed.']
Incoming Data found.
['Fri Sep 16 12:40:05 2016', 'Button 6 pressed.']
10.192.81.112 - - [11/Sep/2016 20:13:39] "GET / HTTP/1.1" 200 2
Incoming Data found.
['Fri Sep 16 12:40:07 2016', 'Button 6 pressed.']
10.192.81.112 - - [11/Sep/2016 20:13:39] "GET / HTTP/1.1" 200 2

```

In the above output you can see 3 instances of **Incoming Data found.** indicating an alert recieved from the Arduino, and 3 API requests recieved.  

# Known Caveats 

