# Automating Deployment of Cisco vWLC on VMWare Fusion

This Ansible playbook can be used to automate deployment of the Cisco Virtual Wireless LAN Controller (vWLC) on VMWare Fusion.  See <a href="http://pseudo.co.de/ansible-cisco-vwlc/" target="_blank">my blog post</a> for an in depth discussion on this playbook.

This has been tested on the following system:

- OS X 10.10.5 (Yosemite)
- VMWare Fusion 7.1.2
- Ansible 1.9.3
- Cisco vWLC version 8.0.120.0

## Prerequisites

- Mac OS X
- <a href="http://www.vmware.com/products/fusion" target="_blank">VMWare Fusion</a> 7.x or higher
- <a href="https://www.vmware.com/support/developer/ovf/" target="_blank">VMWare OVF Tools</a> 4.1 or higher (VMWare account may be required)
- <a href="http://www.ansible.com/" target="_blank">Ansible 1.9.x or higher (`brew install ansible`)
- <a href="https://software.cisco.com/download/release.html?mdfid=284464214&softwareid=280926587&release=8.0.120.0&relind=AVAILABLE&rellifecycle=ED&reltype=latest" target="_blank">Cisco vWLC OVA Image</a> (CCO login required)

This playbook also relies on the Ansible Galaxy <a href="https://github.com/yaegashi/ansible-role-blockinfile" target="_blank">yaegashi.blockinfile module</a>.  You must install this module prior to running the playbook:

```bash
$ ansible-galaxy install yaegashi.blockinfile
- downloading role 'blockinfile', owned by yaegashi
- downloading role from https://github.com/yaegashi/ansible-role-blockinfile/archive/v0.5.tar.gz
- extracting yaegashi.blockinfile to /usr/local/etc/ansible/roles/yaegashi.blockinfile
- yaegashi.blockinfile was installed successfully
```

## Quick Start

First, clone this repository:

```bash
$ git clone https://github.com/cloudhotspot/ansible-cisco-vwlc.git
$ cd ansible-cisco-vwlc
``` 

Next, configure variables for your environment appropriately as described in the <a href="#configuration-variables">Configuration Variables</a> section.  

To run the playbook for a new environment, use the `ansible-playbook site.yml` command.  You will be prompted for your password for sudo access:

```bash
$ ansible-playbook site.yml
SUDO password: *******

PLAY [Check if VM already exists] *********************************************

GATHERING FACTS ***************************************************************
ok: [localhost]
...
...
PLAY RECAP ********************************************************************
localhost                  : ok=36   changed=17   unreachable=0    failed=0
```

This will deploy a virtual machine to the following location:

`{{ vm_destination }}/{{ vm_name }}.vmwarevm/`

For example, if `vm_destination` is **/Users/alice/vms** and `vm_name` is **wlc01** the virtual machine will be deployed to **/Users/alice/vms/wlc01.vmwarevm**.

### Overwriting an Existing Virtual Machine

If an existing virtual machine exists at the target location, by default the playbook will fail with the following error:

```bash
TASK: [Fail if VM path exists] ************************************************
failed: [localhost] => {"failed": true}
msg: VM already exists.  Please set vm_overwrite variable to any value to overwrite the existing VM
``` 

To overwrite a previous deployment, you can set the `vm_overwrite` variable to any value, specifying this either in `vm_vars.yml` or by passing it as an extra variable via the command line (recommended):

`$ ansible-playbook site.yml --extra-vars vm_overwrite=true`

## <a name="configuration-variables"></a>Configuration Variables

There are two files included in the playbook for configuration variables:

- `vm_vars.yml` - configures the OVA source image, virtual machine settings, TFTP and DHCP settings
- `wlc_vars.yml` - configuration settings for the vWLC appliance

### vm_vars.yml

In general, you will need to modify the following variables appropriately for your environment:

```yaml
# Location of the Cisco vWLC OVA image 
ova_source: "/Volumes/Promise RAID/Downloads/AIR-CTVM-K9-8-0-120-0.ova"

# Location of the Cisco vWLC virtual machine that will be created
vm_destination: "/Volumes/Promise RAID/Virtual Machines.localized"

# Name of the Cisco vWLC virtual machine that will be created
vm_name: "wlc01"
```

### wlc_vars.yml

These variables are self explanatory.  At a minimum you should modify the network settings approriately for your environment and choose a different admin password.

```yaml
---
wlc_name: wlc01
wlc_admin_username: admin
wlc_admin_password: Pass1234

wlc_mgmt_ip_address: 192.168.1.6
wlc_mgmt_ip_mask: 255.255.255.0
wlc_mgmt_ip_gateway: 192.168.1.254
wlc_dhcp_server_ip_address: 192.168.1.254 
wlc_virtual_ip_address: 1.1.1.1

wlc_mobility_group_name: "{{ wlc_name }}"
wlc_rf_network_name: "{{ wlc_name }}"
wlc_ntp_server: 64.99.80.30
wlc_ntp_interval: 3600
wlc_ssid: Test SSID
```

## Bring your own vWLC configuration

You can deploy your own vWLC configuration file to the deployed vWLC virtual machine by configuring the `vm_config_file` variable:

`$ ansible-playbook site.yml --extra-vars wlc_config_file=/path/to/ciscowlc.cfg`

You should name your configuration file `ciscowlc.cfg` or use a filename that follows the vWLC AutoInstall conventions (see <a href="http://www.cisco.com/c/en/us/td/docs/wireless/controller/7-2/configuration/guide/cg/cg_gettingstarted.html#wp1144143" target="_blank">how the Cisco vWLC selects a configuration file</a>)

## Cleanup

The playbook installs a temporary DHCP reservation for the vWLC virtual machine which is removed once the vWLC deployment is complete.

The playbook will return the TFTP server run state to its previous state, however the playbook does not revert the previous TFTP configuration file.
