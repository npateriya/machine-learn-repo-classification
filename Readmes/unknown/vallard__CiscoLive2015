# CiscoLive2015

These are Ansible Scripts that deploy a continuous Integration 
Environment on an OpenStack Cloud.  Please see my blog entries 
for more information:

http://benincosa.com/?p=3319

This was also demonstrated at CiscoLive in San Diego in several places:

* The DevNet Zone
* The DevNet Theatre Presentation
* Breakout Session BRKCLD-1003

Video here: 

* https://youtu.be/pLqi4m7b5jk


## Getting Started

This is mostly just used here to show as an example and probably won't just work
right away.  Components like setting up the Docker containers should work as is, but you
may have to change the the ./inventory/hosts to match what you have.    

### Security Component
I've enabled encryption with Ansible Vault.  You can modify or delete this.  
Follow this tutorial here:
http://benincosa.com/?p=3235

To see how we do encryption.  

If you just want to go quick without encryption, do the following:

1. remove:
./vault_passphrase.gpg
./open_the_vault.sh

2. Comment the ./ansible.cfg so it looks like:  

```
  #vault_password_file=open_the_vault.sh
```

3.  Create or rewrite ./roles/gitlab/vars/main.yml 
and define the password for the database: 

```
gitlab_db_password: secret-password
```

# Going Further

Since the containers aren't dynamic here we have a need for a cluster manager.  What do we do about load balancing, etc?  Check out this project from Cisco that addresses this: 

https://github.com/CiscoCloud/microservices-infrastructure



