openconnect
=========

Configure CLI openconnect client with open-vpn
After playbook execution you can connect to VPN simply running vpn_connect or the binary you provided in openconnect_script_name

e.g.

    [vmindru@zetor ~]$ vpn_connect
    Thu Aug 10 13:37:50 2017 TUN/TAP device tun1 opened
    Thu Aug 10 13:37:50 2017 Persist state set to: ON
    POST https://XXXXXXXXXXXXXX.com
    Connected to 001.000.000.000:443
    SSL negotiation with XXXXXXXXXXXXXXXX.com
    Connected to HTTPS on XXXXXXXXXXXXXXXX.com
    XML POST enabled
    Please enter your username and password.
    POST XXXXXXXXXXXXXXXX.com
    XML POST enabled
    Please enter your username and password.
    Password:
    POST https://XXXXXXXXXXXXXXXX.com
    Got CONNECT response: HTTP/1.1 200 OK
    CSTP connected. DPD 30, Keepalive 20
    Connected as 172.16.0.228, using SSL
    Established DTLS connection (using GnuTLS). 




Requirements
------------

this has been tested on Fedora only

Role Variables
--------------

    openconnect_script_template: "vpn-client.sh.j2"
    openconnect_script_name: "vpn_connect"
    openconnect_script_dest: "/usr/bin/{{ openconnect_script_name }}"
    openconnect_vpn_group: "none"
    openconnect_vpn_user: "none"
    openconnect_vpn_url: "none"
    openconnect_vpnc_script: "/etc/vpnc/vpnc-script" #this is default value for RHEL boxes, might need adjustments on playbook level for debian

Example Playbook
----------------

you can use the playbook provided with the role 

    [vmindru@zetor openconnect]$ ansible-playbook -i localhost, demo_playbook.yml  -c local -b -K                                                                                                                      
    SUDO password:                                                                                                                                                                                                     
    provide VPN group: RocketFuel                                                                                                                                                                                      
    provide VPN user: vmindru                                                                                                                                                                                          
    provide anyconnect vpn URL: https://some_fqdn           
    [vmindru@zetor openconnect]$

or create your own similar playbook

		- hosts: localhost
	  connection: local
	  vars_prompt:
	    - name: openconnect_vpn_group
	      prompt: "provide VPN group"
	      private: no
	
	    - name: openconnect_vpn_user
	      prompt: "provide VPN user"
	      private: no
	
	    - name: openconnect_vpn_url
	      prompt: "provide anyconnect vpn URL"
	      private: no
	  roles:
	  - { role: openconnect }



License
-------

GPLv2

Author Information
------------------
Veaceslav Mindru m1ndruvXgma1l.c0m
