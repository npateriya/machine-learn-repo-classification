Step 1: Git clone the playbooks

Git must be cloned on the NFVI management (build) node. Any directory can be used.

Step 2: Change inventory/hosts file

TODO: this should be changed to dynamic inventory generation based on NFVI setup.yaml file

Example:

[root@nfvi-mgmt nfvi-monitorix]# more inventory/hosts
[local]
127.0.0.1 ansible_connection=local ansible_python_interpreter="/usr/bin/python"

[control]
control-server-1
control-server-2
control-server-3

[storage]
storage-server-1
storage-server-2
storage-server-3

[compute]
compute-server-3
compute-server-4
compute-server-5
compute-server-6
compute-server-7
compute-server-8

Step 3: Change configuration yaml file

Example:

# more config/setup.yaml
title: "Cisco NFVI C-POD in Paris"
root_dir: "/root/demo/monitorix"
port: 19124

Note: 'root_dir' is the target directory where monitoring scripts and files will be copied to. This must not be the directory where git is cloned to.

Step 4: Run installation playbook

ansible-playbook playbooks/install-monitorix.yaml

This playbook will install all files and startup scripts on all NFVI nodes (configured in inventory/hosts).

Since the only Docker repository allowed on the NFVI nodes is the local Docker repository on the management node, Docker image is saved into tarball and added from it.

Step 5: Start Monitorix playbook

ansible-playbook playbooks/start-monitorix.yaml

Step 6: (Optional) NAT rules to access Monitorix web page

An idea to use multihost feature of Monitorix has one big disadvantage that you cannot get all graphs for node other than localhost (build node).
http://www.monitorix.org/faq.html#Q111

Example:
ip nat inside source static tcp 192.168.27.102 19124 10.60.21.144 20000 extendable
ip nat inside source static tcp 192.168.27.100 19124 10.60.21.144 20001 extendable
ip nat inside source static tcp 192.168.27.101 19124 10.60.21.144 20002 extendable
ip nat inside source static tcp 192.168.27.105 19124 10.60.21.144 20003 extendable
ip nat inside source static tcp 192.168.27.103 19124 10.60.21.144 20004 extendable
ip nat inside source static tcp 192.168.27.104 19124 10.60.21.144 20005 extendable
ip nat inside source static tcp 192.168.27.106 19124 10.60.21.144 20006 extendable
ip nat inside source static tcp 192.168.27.105 19124 10.60.21.144 20007 extendable
ip nat inside source static tcp 192.168.27.108 19124 10.60.21.144 20008 extendable
ip nat inside source static tcp 192.168.27.109 19124 10.60.21.144 20009 extendable
ip nat inside source static tcp 192.168.27.107 19124 10.60.21.144 20010 extendable
ip nat inside source static tcp 192.168.27.111 19124 10.60.21.144 20011 extendable

