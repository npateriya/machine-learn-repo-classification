# metacloud_ansible_hadoop
Deploying Hadoop with Ansible on Cisco Metacloud OpenStack

Step 1:
Run Bash script to start ansible playbooks 

~$ . hadoopplaybook

It invokes these playbooks
hadoopvm.yaml  #Creates volumes, kepairs, instances, and copies assigned IPs to host.txt and other txt files to create ansible inventory and  /etc/hosts for the virtual machines

installambari-master.yaml #copy host files, mnt volume, install amabari server and agent

installambari-data1.yaml #copy host files, mnt volume, install amabari agent

installambari-data2.yaml #copy host files, mnt volume, install amabari agen

blueprintcluster.yaml # curl API commands to register and deploy hadoop cluster blueprint into Ambari

Access Ambari with ambari-master IP address:8080  
admin/admin


