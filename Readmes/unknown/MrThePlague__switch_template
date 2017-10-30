# Switch Template

Quick and dirty Ansible playbook to create switch configurations for Cisco IOS/IOS-XE.

## Instructions

1. Clone repo
2. Edit `group_vars/all.yml` with global variables, such as syslog servers, NTP servers, etc.
3. Edit `roles/access/vars/main.yml` with switch specific variables, such as IP addresses, VLANs, etc.
4. Execute `$ ansible-playbook switch_template.yml`

### To-do:

* Make this README not suck
* Create roles roles for specific switch models
* Document the jinja2 template file
* Nest variables better
* Create loop for user access ports
