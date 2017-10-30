The entire installation should be completed from a contol node or your laptop if you wish.  Be aware your control node will require various packages that can be referenced at https://github.com/CiscoCloud/mantl.


1.) Update the inventory file to include your servers.


2.) Copy ssh-key from your control server to all nodes.  
Example:
ssh-copy-id steveno@c7-2


3.) copy host file from control server to all other hosts.  The host file should cover your entire inventory.
    ansible all -i local_inventory.yml -u steveno --sudo --ask-sudo-pass -m copy -a "src=/etc/hosts dest=/etc/hosts"

4.) Run this install on control node:
	sudo -H pip install -r requirements.txt

5.)
Command to run install:

python install_mantl.py
