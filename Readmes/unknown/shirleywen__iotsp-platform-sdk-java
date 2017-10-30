# Cisco IOT Software Platform SDK in Java

## Description

A software development kit to enable you to build Java applications on the IOT Software Platform.

## Getting Started

1. Download or clone Cisco IOT Software Platform SDK in Java
2. Install JDK, recommend version 8
3. Dependencies are managed by gradle. Run the following command to download dependencies and build the SDK:
    on Windows:
      gradlew.bat build
    on Mac or unix system:
      ./gradlew build
4. Upon successful build, jar file iotsp-java-sdk-1.0.0-all.jar will be generated
5. Include the jar file iotsp-java-sdk-1.0.0-all.jar in your Java project


Once SDK is downloaded or cloned, you will see a project structure like below:
```shell

├── src
|   ├── main
|       └── java            
|       |   └── com                
|       |       └── cisco
|       |           └── iotsp
|       |               ├── client                          # Libary packages for a cloud app to import and use to communicate with IoTSP services
|       |               |   └── accounts                    # Libary package to communicate with IoTSP accounts service
|       |               |   └── claims                      # Libary package to communicate with IoTSP claims service
|       |               |   └── http_device_connector       # Libary package to communicate with IoTSP http_device_connector service
|       |               |   └── observations                # Libary package to communicate with IoTSP observations service
|       |               |   └── presence                    # Libary package to communicate with IoTSP presence service
|       |               |   └── registration                # Libary package to communicate with IoTSP registration service
|       |               |   └── schemas                     # Libary package to communicate with IoTSP schemas service
|       |               |   └── thing_credentials           # Libary package to communicate with IoTSP thing_credentials service
|       |               |   └── things                      # Libary package to communicate with IoTSP things service
|       |               |   └── users                       # Libary package to communicate with IoTSP users service
|       |               ├── helper                          # helper package that a cloud app can import
|       |               |   └── AccountHelper.java          # helper class to create an account with IoTSP
|       |               |   └── AuthenticationHelper.java   # helper class to generate access token used to communicate with IoTSP
|       |               |   └── ClientHelper.java           # helper class that creates https client to communicate with IoTSP
|       |               |   └── JsonHelper.java             # helper class converst Json to and from objects
|       |               |   └── ServiceApiHelper.java       # helper class to generate proxy clients (API classes )for a cloud app to communicate with IoTSP services  
|       |               ├── sample                          # sample code on how cloud app can communicate with various IoTSP's services
|       |               |   └── workflow                    # sample code on how to implement several use cases
|       |               └── test                            # testing app that a cloud app developer may run to step into each sample code
|       |                   └── TestAll.java                # testing app that test all the sample code
|       └── resources
|          └── file                                         # sample data files used by sample code 
├── gradlew                                                 # gradle wrapper to be executed in Mac or Unix system to build SDK
└── gradlew.bat                                             # gradle wrapper to be executed in Windows platform to build SDK
```
***Note: when run testing app, it creates accounts, users, schemas, devices on IoTSP, then delete all of them if testing app is ran successfully.
