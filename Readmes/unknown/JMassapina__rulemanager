# rulemanager.rb

## What is this thing?

A Ruby-based daemon to populate an object-group on all your Cisco ASAs, so that it contains all of your Amazon EC2 instances. It works over multiple AWS accounts and regions, and synchronizes multiple Cisco devices. It also includes an SQS consumer to allow new instances to be added instantly.

## How does it work?

This tool accepts a configuration file containing AWS access keys, and generates a list of IP addresses for all accessible EC2 hosts in all given regions. It then connects to any number of Cisco ASA devices, and synchronizes the list of IP addresses as members of a dyanmic object-group, which can be used with access-lists to allow access from EC2 hosts into a corporate network.

A simple cache is maintained, so that a full sync will only run if the list of IP addresses of your active nodes has changed.

It also includes an SQS consumer, which will subscribe to any given SQS queue, and can then act immediately as new instances are brought up.

## Who developed this?
This tool was originally developed by James Denness (scarybot) and the network operations team at Mind Candy.

## Configuration

### Server
There are two files you'll need to care about. The first is the general config file (rulemanager.conf.yml). A template for this is included in this repo. The settings available in there should be reasonably self-explanatory. The second file is your AWS account information. The path to this file is specified in the main rulemanager.conf.yml, and it should be of the format:

    [aws-account-name@my-host.com]
    aws_access_key_id = YOUR_ACCESS_KEY_ID
    aws_secret_access_key = YOUR_VERY_SECRET_ACCESS_KEY
    regions = ['eu-west-1', 'us-east-1']

### Cisco
You'll need to create an object-group on any ASA you'd like to manage. For example:

    asa> enable
    asa# configure terminal
    asa(config)# object-group network dynamic-ec2-hosts

It's also necessary to allow SSH access from whichever host(s) or network(s) you'll be running rulemanager on:

    asa(config)# ssh 192.168.0.50 255.255.255.0 inside

Don't forget to allow SSH connections through your access-lists also.

### Amazon
Create an IAM user in AWS with an access policy similar to the following:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*",
        "sqs:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
```

You may wish to create two separate users; one to be used by rulemanager (requiring SQS ListQueues/RecieveMessage and EC2 describe privileges), and another to be used for submitting 'host up' notifications to the queue (requiring SQS:*).

Then, create an SQS queue. This will be used to notify rulemanager of newly spun-up hosts. Copy the name of the queue into the config file.

### Clients
You don't actually need to do anything on the client; you could choose to rely on rulemanager's full-sync capability every few minutes, in order to bring your firewalls up-to-date with new machines. If your new machines need immediate access to remote resources, then you're going to want to make use of the SQS functionality in order to have the firewalls updated as soon as a new node is brought to life. The simplest way to do this, is to have your instance userdata take care of it. One way is to store a script on S3, and have the userdata retrieve and run it on first boot. Here's an exomple python script:

```python
#!/usr/bin/env python
import requests
import boto
import json
import boto.ec2
import boto.s3
import boto.sqs
import time
import os
from boto.sqs.message import Message

url = 'http://169.254.169.254/latest/meta-data/'

region_name = requests.get(url+'placement/availability-zone').text[:-1]
region_full = requests.get(url+'placement/availability-zone').text
instance_id = requests.get(url+'instance-id').text
public_hostname = requests.get(url+'public-hostname').text
mac = requests.get(url+'network/interfaces/macs/').text
owner = requests.get(url+'network/interfaces/macs/' + mac + '/owner-id').text

conn = boto.sqs.connect_to_region('eu-west-1')
q = conn.get_queue(queue_name='autoscaling', owner_acct_id='YOUR_AWS_ACCOUNT_NUMBER')
m = Message()
m.set_body(json.dumps({ "instance_id": instance_id, "hostname": public_hostname, "state": 'up', "region": region_full, "owner": owner }))
status = q.write(m)
```

This script is also available in the conf/ directory of this repository. Be sure to replace your AWS owner account ID, or if you're only working in a single AWS account, you can remove that parameter. Essentially, you need to add a message to your queue with the following JSON:

```json
{
    "instance_id": instance_id,
    "hostname": hostname,
    "state": "up",
    "region": region_including_availability_zone,
    "owner": this_instance_owner_aws_account_number
}
