
--------------------------------------------------------------------------------
   device_sdk: (Supports v1 and V2 version of the API's)
--------------------------------------------------------------------------------
Device-sdk project builds as a C library and applications that want to use this library
should include the sdk header files and link against the library .so file.
This project contains the device_sdk library source files and sample application source files.

The device_sdk library provides the application with an interface to talk to IOTSP.
In short, device_sdk abstracts all the apis's exposed by IOTSP.
The sample application provides an example on how to exercise the API's exposed by
the device_sdk library.

There are two version of api sets exposed by the device_sdk library
  - V2 newer version API's(Recommended )
  - V1 older version API's

sampleAppV2 - shows how to use the V2 Api's
  - application code just needs to call the send data API
  - device registration, access token request and data posting is all taken care in one single API

sampleApp - Shows how to use the V1 APi's
 - Application code has more control as it exercises individual API's
  - to initialize the device sdk with device details
  - to set the logging level and log file information
  - to get the device registered to the registration service.
  - to acquire an Access Token
  - to send device data to IOTSP

--------------------------------------------------------------------------------
   Project (main) Files & Directory :
--------------------------------------------------------------------------------
1.) api :
  This directory has the io_device_sdk.h header file. The header file that the application
  code has to include.

2. ) config :
  This directory holds the sdk_config.json file. This has the config parms that
  can be passed to the device_sdk library.
  This is the file where the IOTSP server ip address is set, device data connector type is set etc.

3.) docs :
  This directory has the device_sdk user guide in pdf form.

4.) install :
  This directory which holds all the dependency libraries and this is where the
  devic_sdk libary will be installed.

5.) linux :
  This directory holds the device_sdk source files.
  - iot_device_sdk_init.c        [ initialization ]
  - iot_device_sdk_config.c      [ sdk config load]
  - iot_device_sdk_logging.c     [ sdk logging related]
  - iot_device_sdk_serialize.c   [ data serialization ]
  - iot_device_sdk_json.c        [sdk json abstraction api's]
  - iot_device_sdk_curl.c        [ libcurl wrapper api's ]
  - iot_device_sdk_rest.c        [ rest wrapper api's ]
  - iot_device_sdk_default_msg.c [ sdk's default json messages ]
  - iot_device_sdk_data_send.c     [ sdk data send API's]
  - iot_device_sdk_helper.c        [ sleep api]
  - iot_device_sdk_mqtt_client.c   [ mqtt apis]
  - iot_device_sdk_registration.c  [ sdk registration API's ]
  - iot_device_sdk_storage.c       [ sdk storage api's to store device conf details]
  - iot_device_sdk_storage_cache.c [ sdk storage cache api's]
  - iot_device_sdk_utils.c         [ sdk list implementation to store objects]
  - iot_device_sdk_access_token.c  [ sdk access token api's]
  - iot_device_sdk_mqtt_client.c   [ sdk mqtt related code]

6.) sampleApps :
 This directory contains all the sample reference Applications.

6.1) sampleApp:
 This sample App shows how to use sdk v1 API's.
 Path: iot_device_sdk\sampleApps\sampleApp
 The sample contains
 - sample_main.c         [ basic application loop with calls to init, initiate registration
                          serialize and Send device Data ]
 - device_registration.c [ registration sequence with the servers, error handling,retries etc ]
 - helper.c              [ helper functions/code ]

6.2) sampleAppV2:
 This sample App shows how to use sdk v2 API's.
 (data send V2 API - is the only call that the application has to make -
   internally the sdk will register the device and setup the mqtt session and
   then posts the device data)

 Path: iot_device_sdk\sampleApps\sampleAppV2
 The sample contains
 - sample_main.c         [ basic application loop with calls to data send V2 API ]
 - helper.c              [ helper functions/code ]

6.3) sampleAppSerial :
 This sample App shows how to use sdk v2 API's to send sensor data from serial port.
 (data send V2 API - is the only call that the application has to make -
   internally the sdk will register the device and setup the mqtt session and
   then posts the device data)

 Path: iot_device_sdk\sampleApps\sampleAppSerial
 The sample contains
 - sample_main.c         [ basic application loop with calls to data send V2 API ]
 - helper.c              [ helper functions/code ]


--------------------------------------------------------------------------------
  Build Steps:
--------------------------------------------------------------------------------
Dependency libraries
 - libcurl (sdk lbrary verfied aganist 7.48.0)
 - jansson-2.7
 - paho mqtt c (1.0.3)

Dependency libraries are captured in a separate git project.
(IOTSP/device-sdk-c-dep-libs  : https://cto-github.cisco.com/IOTSP/device-sdk-c-dep-libs )
Build the above project and copy or install the lib/bin/install folders to
device-sdk/install directory. So that the dependency libraries header and .so files are available to
device_sdk.

Build step:
  - cd to main device-sdk directory
  - do #> make all
  - This will buid lib_iot_device_sdk.so device sdk library
    and all the sample applications

 SampleApp Run Step (example shows how to use sdk v1 APIS'):
  - cd to the sampleApps/sampleApp directory
  - #> LD_LIBRARY_PATH=../../install/lib ./sampleApp -d deviceInfo.txt
  (one can use deviceInfo.txt file to feed the device details info to the application )

 SampleAppV2 Run Step (example shows how to use sdk v2 APIS'):
  - cd to the sampleApps/sampleAppV2 directory
  - #> LD_LIBRARY_PATH=../../install/lib ./sampleAppV2 -d deviceInfo.txt
  (one can use deviceInfo.txt file to feed the device details info to the application )
