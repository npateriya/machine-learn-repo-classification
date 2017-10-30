# Cisco IOT Software Platform SDK in Python

## Description

A software development kit to enable you to build Python applications on the IOT Software Platform.

## Getting Started

1. Download or clone Cisco IOT Software Platform SDK in Python
2. Install Python 2.7
3. Run the following command to install the required libraries listed in requirements.txt: </br>
    **_pip install -r requirements.txt_**
4. Include SDK source code, choose the IoTSP services, and import the corresponding packages under com/cisco/iotsp/sdk into your Python application. Then distribute SDK with your application.</br>
or</br>
5. Distribute and install SDK by itself: </br>
    a. Run the following command to create a source distribution package for SDK: </br>
        **_python setup.py sdist_** </br>
    b. Copy the iotsp-sdk-python-1.0.0.zip file from dist directory, to the target platform, and unzip the file. </br>
    c. Run the following command to install SDK: </br>
        **_python setup.py install_**

Once downloaded or cloned, you will see a project structure like below:
```shell

├── com
|   ├── cisco
|      └── iotsp
|          ├── helper                          # helper package that a cloud app can import
|          |   └── account_helper.py           # helper class to create an account with IoTSP
|          |   └── authentication_helper.py    # helper class togenerate access token used to communicate with IoTSP
|          ├── sample                          # sample code on how cloud app can communicate with IoTSP's services
|          |   └── workflow                    # sample code on how to implement several use cases
|          |   └── data                        # sample data files used by sample code
|          ├── sdk                             # Libary packages for a cloud app to import and use to communicate with IoTSP services
|          |   └── accounts                    # Libary package to communicate with IoTSP accounts service
|          |   └── claims                      # Libary package to communicate with IoTSP claims service
|          |   └── http_device_connector       # Libary package to communicate with IoTSP http_device_connector service
|          |   └── observations                # Libary package to communicate with IoTSP observations service
|          |   └── presence                    # Libary package to communicate with IoTSP presence service
|          |   └── registration                # Libary package to communicate with IoTSP registration service
|          |   └── schemas                     # Libary package to communicate with IoTSP schemas service
|          |   └── thing_credentials           # Libary package to communicate with IoTSP thing_credentials service
|          |   └── things                      # Libary package to communicate with IoTSP things service
|          |   └── users                       # Libary package to communicate with IoTSP users service
|          └── test                            # testing app that a cloud app developer may run to step into each sample code
|          |    └── test_all.py                # testing app that test all the sample code
|          |    └── test_api.py                # testing class that test sample code under sample\
|          |    └── test_workflow.py           # testing calss that test sample code under sample\workflow
|          └── services.properties             # contains IoTSP cluster IP address that test_all.py uses to run testing app
└── requirements.txt                           # contains required libraries
```

Run the included test_all.py application using one of these methods:

1. Run test_all.py inside an IDE, such as PyCharm. <br/>
or </br>
2. Build the SDK on your target platform using step 5 (above), and then enter the following command in the CLI: </br>
    **_python test_all.py_**

Note: The python test_all.py application creates accounts, users, schemas, and devices on IoTSP, and then deletes them if the app executes successfully.
