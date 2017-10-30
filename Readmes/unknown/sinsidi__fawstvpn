# fawstvpn

Fawstvpn is a python script that generates and applies the configuration for an AWS VPN in a Cisco IOS router.
## Installation
### Debian / Ubuntu
```sh
$ apt-get install python3-paramiko python3-boto3
```
### Generic Linux
```sh
$ pip3 paramiko
$ pip3 install boto3
```
### Command line options
```sh
--config_file or -c
    Default='./fawstvpn.cfg'. Help='Configuration file'
--vpn_connection_id or -v
    Default='Mandatory parameter, no default'. Help='AWS VPN Connection ID'
--dry_run or -d
    Default='False'. Help='Do not apply changes, just screen printing'
--force or -f
    Default='False'. Help='Do not apply changes, just screen printing'
--delete or -r
    Default='False'. Help='Remove an existing configuration instead of create a new one'
``` 
### Configuration files
There is a main section in the configuration file that defines the global parameter for the application.
```sh
[Main]
# AWS Region where the VPN is located
aws_region:						eu-west-1 
# Logging detail
log_level: 						debug
# The template to create the VPN. Local file 
create_template: 				file:path/generic_cisco_create.tmpl
# ...or S3 is supported
create_template: 				s3:bucket-name/path/generic_cisco_create.tmpl
# The template to create the script with the command to delete the VPN. Local file
delete_template: 				file:path/generic_cisco_delete.tmpl
# ...or S3 is supported
delete_template: 				s3:bucket-name/path/generic_cisco_delete.tmpl
# This defines if the script to delete the configuration must be created
delete_script: 			true
# The script can be saved in a local file
delete_script_folder:			file:/tmp/
# ... or in a S3 bucket.
delete_script_folder:			s3:bucket-name/path
# And the aws credentials. If not defined, get ones from ~/.aws/credentials
aws_access_key_id:				
aws_secret_access_key:			
```
The configuration file must include a section for those VPNs that should be managed by this script. The name of the section must be the Customer Gateway Id that represents the physical router in the AWS VPC.
```sh
[cgw-12345678]
# Some parameters of the main section can be overridden in each customer gateway
create_template:		file:generic_cisco_create.tmpl
delete_template: 		file:generic_cisco_delete.tmpl
delete_script: 	true
# These parameters are required to establish the SSH connection with the physical router
hostname: 				192.168.100.2
username: 				admin
password: 				admin
# The following parameters gives the chance to run some parameters before and after the command to retrieve the configuration prior to apply the VPN configuration.
pre_config_commands:	enable,admin
post_config_commands:	copy running-config startup-config,startup-config
# And the command that must be run to get the configuration
get_config_commands:	terminal length 0,show run
# And, finally, some labels can be defined to be sustituted in the configuration templates
advertised_network: 	192.168.90.0 mask 255.255.255.0
isakmp_policy_number:	200 
uplink_address: 		FastEthernet0/0
transform_set_name:		aws-transform-set
ipsec_profile_name:		aws-ipsec-profile
```
### Configuration templates
In order to generate a valid configuration from the configuration templates, some keys must be used:
* **Config parameters** (Ex. {config:isakmp_policy_number}). They are defined in the configuration file and will be replaced by the specified value.
* **XML parameters** (Ex. {xml:ike/pre_shared_key}). They point to the values in the XML file where AWS stores the VPN configuration. It can be retrieve using the [CLI/API](http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpn-connections.html) to check the existing keys.
* **Conflict numbers** (Ex. {conflict-number:interface Tunnel}). Defined to organize the numbering of the interfaces. The example represents that the script will find a free number for "interface Tunnel".  
* **Repetition key** (Ex. {repetition}) As the AWS VPN configuration has several VPNs configuration, this represents the iteration. There should be a couple of iterations, so it will be 0 or 1.
* **VPN Connection Id key** (Ex. {vpn_connection_id}). The ID of the VPN that the script is configuring. If the scripts already detects this ID in the current configuration, it assumes that the configuration for this VPN was already applied.