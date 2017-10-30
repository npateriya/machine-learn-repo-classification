# ansible-vwlc-role

An Ansible role for automating deployment of the Cisco Virtual Wireless Controller (vWLC) on VMWare Fusion.  

The role configures Mac OS X TFTP and VMWare DHCP to enable the AutoInstall feature of the vWLC.

##Requirements

- Mac OS X
- <a href="http://www.vmware.com/products/fusion" target="_blank">VMWare Fusion</a> 7.x or higher
- <a href="https://www.vmware.com/support/developer/ovf/" target="_blank">VMWare OVF Tools</a> 4.1 or higher (VMWare account may be required)
- <a href="https://software.cisco.com/download/release.html?mdfid=284464214&softwareid=280926587&release=8.0.120.0&relind=AVAILABLE&rellifecycle=ED&reltype=latest" target="_blank">Cisco vWLC OVA Image</a> (CCO login required)

##Role Variables

You must specify the following variables in your playbook:

```yaml
# Location of the Cisco vWLC 1000V OVA image 
wlc_ova_source: "/path/to/ova/source"

# Root folder where the Cisco vWLC virtual machine will be created
wlc_vm_root: "/path/to/root"
```

The vWLC virtual machine will be deployed to the following location:

`{{ wlc_vm_root }}/{{ wlc_vm_name }}.vmwarevm/`

For example, if `wlc_vm_root` is **/Users/alice/guests** and `wlc_vm_name` is **wlc01** the virtual machine will be deployed to **/Users/alice/guests/wlc01.vmwarevm**.

If the virtual machine already exists, by default the role will fail.  To overwrite the existing virtual machine, the following variable must be set (to any value):

`wlc_vm_overwrite: yes`

###Default Role Variables

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

Dependencies
------------

This role relies on the Ansible Galaxy <a href="https://github.com/yaegashi/ansible-role-blockinfile" target="_blank">yaegashi.blockinfile module</a>.  Installing this role will automatically install this module.

Example Playbook
----------------

This playbook is designed to run locally on local OS X host so you should configure any play that uses this role with `hosts: localhost` and `connection: local`:

    - hosts: localhost
      connection: local
      roles:
         - { role: mixja.vwlc, wlc_vm_overwrite: true, wlc_ova_source: /path/to/ova/source, wlc_ova_root: /path/to/vm/root  }

A sample playbook is provided at <a href="https://github.com/cloudhotspot/ansible-vwlc-playbook">https://github.com/cloudhotspot/ansible-vwlc-playbook</a>

License
-------

BSD

Author Information
------------------

Created by Justin Menga - see http://pseudo.co.de
