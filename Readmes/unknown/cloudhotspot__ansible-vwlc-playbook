# Sample Playbook for Automating Deployment of Cisco vWLC on VMWare Fusion

This is an example playbook that can be used to automate deployment of the Cisco Virtual Wireless LAN Controller (vWLC) on VMWare Fusion using Ansible.  See <a href="http://pseudo.co.de/ansible-cisco-vwlc/" target="_blank">my blog post</a> for an in depth discussion on this playbook.

This has been tested on the following system:

- OS X 10.10.5 (Yosemite)/OS X 10.11 (El Capitan)
- VMWare Fusion 7.1.2
- Ansible 1.9.3
- Cisco vWLC version 8.0.120.0

## Prerequisites

- Mac OS X
- <a href="http://www.vmware.com/products/fusion" target="_blank">VMWare Fusion</a> 7.x or higher
- <a href="https://www.vmware.com/support/developer/ovf/" target="_blank">VMWare OVF Tools</a> 4.1 or higher (VMWare account may be required)
- <a href="http://www.ansible.com/" target="_blank">Ansible 1.9.x or higher (`brew install ansible`)
- <a href="https://software.cisco.com/download/release.html?mdfid=284464214&softwareid=280926587&release=8.0.120.0&relind=AVAILABLE&rellifecycle=ED&reltype=latest" target="_blank">Cisco vWLC OVA Image</a> (CCO login required)

And of course the `mixja.vwlc` role:
```bash
$ ansible-galaxy install mixja.vwlc
- downloading role 'vwlc', owned by mixja
- downloading role from https://github.com/mixja/ansible-vwlc-role/archive/v0.5.tar.gz
- extracting mixja.vwlc to /usr/local/etc/ansible/roles/mixja.vwlc
- mixja.vwlc was installed successfully
```


This playbook also relies on the Ansible Galaxy <a href="https://github.com/yaegashi/ansible-role-blockinfile" target="_blank">yaegashi.blockinfile module</a> which is automatically installed when you install the `mixja.vwlc` role.

## Quick Start

First, clone this repository:

```bash
$ git clone https://github.com/cloudhotspot/ansible-vwlc-playbook.git
$ cd ansible-vwlc-playbook
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

`{{ wlc_vm_destination }}/{{ wlc_vm_name }}.vmwarevm/`

For example, if `wlc_vm_destination` is **/Users/alice/vms** and `wlc_vm_name` is **wlc01** the virtual machine will be deployed to **/Users/alice/vms/wlc01.vmwarevm**.

### Overwriting an Existing Virtual Machine

If an existing virtual machine exists at the target location, by default the playbook will fail with the following error:

```bash
TASK: [Fail if VM path exists] ************************************************
failed: [localhost] => {"failed": true}
msg: VM already exists.  Please set wlc_vm_overwrite variable to any value to overwrite the existing VM
``` 

To overwrite a previous deployment, you can set the `wlc_vm_overwrite` variable to any value, specifying this either in your playbook variables or by passing it as an extra variable via the command line (recommended):

`$ ansible-playbook site.yml --extra-vars wlc_vm_overwrite=true`

## <a name="configuration-variables"></a>Configuration Variables

The following variables are available to configure:
  
```yaml
# Name of the Cisco vWLC virtual machine that will be created
wlc_vm_name: "wlc01"

# Last Octet of IP address assigned to vWLC service port.  This value should be between 3 and 127.
wlc_vm_svc_ip_octet: "121"

# Keep the DHCP reservation used for provisioning
wlc_vm_persist_dhcp_reservation: yes

# TFTP settings
wlc_tftp_path: /Users/Shared/tftp

# Location of TFTP plist file
wlc_tftp_plist: /Users/Shared/tftp.plist

# WLC configuration settings
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
## Service Port Configuration

The default intended use case is to configure Ethernet0 (the vWLC service port) as attached to the `vmnet8` interface on the host.  This playbook configures a DHCP reservation based upon `vmnet8` interface IP address and uses the `wlc_vm_svc_ip_octet` value to configure the desired last octet of the service port IP address.

By default this playbook also provisions the service port IP address and `wlc_vm_name` value into the local machines `/etc/hosts` file.

For example, if the `vmnet8` interface subnet is **192.168.100.0/24** (this will vary for each VMWare Fusion environment) and the `wlc_vm_svc_ip_octet` value is **121**, then the service port IP address assigned to the vWLC virtual machine will be **192.168.100.121**.  

If you only wish to temporarily use the dynamic service port IP address (e.g. because you are overwriting the management interface with your own configuration using a `wlc_config_file`), you should remove the DHCP reservation after provisioning by setting the `wlc_vm_persist_dhcp_reservation` variable to **no**.  Setting this variable to **no** will remove the DHCP reservation and also will not provision the management IP address and VM name into the local `/etc/hosts` file.

## Bring your own vWLC configuration

You can deploy your own vWLC configuration file to the deployed vWLC virtual machine by configuring the `wlc_config_file` variable:

`$ ansible-playbook site.yml --extra-vars wlc_config_file=/path/to/ciscowlc.cfg`
