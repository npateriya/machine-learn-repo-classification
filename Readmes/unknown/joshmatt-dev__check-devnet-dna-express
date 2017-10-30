# check-devnet-dna-express
This is a a collection of scripts to check user workstation preparation for Cisco DevNet DNA Express events.

## Usage
### Windows
- Clone this repo into your desired working directory.
- Edit checkDevNet.py and add your Spark Token to `SPARK_TOKEN = ""` (i.e. `SPARK_TOKEN = "abcde12345"`)
- From cmd run `checkDevNet.bat <virt_env_name>`.
- A folder named `virt_env_name` will be created and be used as your working directory for the Cisco DevNet DNA Express event. 

### OSX & Linux
- Clone this repo into your desired working directory.
- Edit checkDevNet.py and add your Spark Token to `SPARK_TOKEN = ""` (i.e. `SPARK_TOKEN = "abcde12345"`)
- From terminal run `source checkDevNet.sh <virt_env_name>`
- A folder named `virt_env_name` will be created and be used as your working directory for the Cisco DevNet DNA Express event. 

## What it does...
1. Checks system for installation of at least Python 3.5
2. Checks system for installation of pip for Python 3
3. Checks system for installation of Python Virtualenv Library
4. Detects or creates Python Virtual Environment
5. Installs Python libraries necessary for DevNet Express event 
6. Creates a Cisco Spark Room and posts a messages to it using Cisco Spark APIs
7. Checks Git installation and clones or pulls updates for https://github.com/CiscoDevNet/devnet-express-code-samples.git repository
