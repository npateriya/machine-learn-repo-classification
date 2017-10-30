## A playbook to load the Ubuntu/Debian agents used with ciscozeus

This is a simple Ansible play that will load the Fluent, collectd, and rsyslog configs as needed to interact with CiscoZeus. To use them, update the inventory file to include the hostnames (and ssh default user) needed to access your Ubuntu or Debian based macines. In addition, either update the zeus_env file, and source the contents, or in some other fashion set a USERNAME and TOKEN environment variable so that the td-agent configuration can be properly set up on the machine.

Then simply run the play as in:

    ansible-playbook -i inventory play.yml

You should shortly have a machine that will send it's current usage data to CiscoZeus!

More on ciscozeus here: http://ciscozeus.io
