# cisco-ansible
managing cisco ios (version 15 ) with  ansible

ansible --version
ansible 2.3.2.0

Tested with Cisco IOS SW L2 and L3 (version 15 only)

At the Ansible server,generate ssh key and then extract the rsa.pub from the below command
root@ubuntu-VirtualBox:# fold -b -w 72 /root/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDc44198a7tXhZUmyRRn8Bkl+Gak/zlrks8
QqZMI8rAq0YAlJWoiw0nQMLobhP+0q/kjs7rbFSCbKtzSxTyGW9uDWD02ru0pk6Lr0ND3m+K
06T3HNu9Tfdz1lmyaZ4fIxbjDCNEup3/eVteWVogJONLcBpt6oeUdfFyYOp7k0TnKQ== roo
t@ubuntu-VirtualBox


Router1(config)#ip ssh pubkey-chain 
Router1(conf-ssh-pubkey)#username sourav
Router1(conf-ssh-pubkey-user)#key-string
Router1(conf-ssh-pubkey-data)#$BAAAAgQDc44198a7tXhZUmyRRn8Bkl+Gak/zlrks8
Router1(conf-ssh-pubkey-data)#$kjs7rbFSCbKtzSxTyGW9uDWD02ru0pk6Lr0ND3m+K
Router1(conf-ssh-pubkey-data)#$/eVteWVogJONLcBpt6oeUdfFyYOp7k0TnKQ== roo
Router1(conf-ssh-pubkey-data)#t@ubuntu-VirtualBox
Router1(conf-ssh-pubkey-data)#exit
Router1(conf-ssh-pubkey-user)#exit
Router1(conf-ssh-pubkey)#exit
Router1(config)#exit


Once,ssh setup is done ,it's very easy to use ansible for managing cisco ios.

