Simple IPsec health checker for Cisco router
============================================

It just logins into the Cisco router you specified in the argument.
And, it takes a set of IP addresses which are assigned to the peer
for the IPsec tunnel.
Then, it tries to send some ping packets to each of them.

The default timeout is 15 seconds for each ping.
Please note that it takes 15 seconds multiplied
by the number of peers at most to check the all peers.

It only supports SSH login.

## REQUIREMENT

- only python2.7 is tested.  python3 is not tested.
- [paramiko](http://www.paramiko.org/), a python module for ssh.  The instruction to install paramiko is [here](http://www.paramiko.org/installing.html).
- known_hosts

You have to login into the server once so as to create
known_hosts before you use this script.

## USAGE

    ipsec-health-check.py [-h] -s SERVER [-u USERNAME] [-p PASSWORD]
                                 [--syslog] [--sshport SSH_PORT]
                                 [--sshmode SSH_MODE] [--timeout TIMEOUT]
                                 [--tagtype {rsa,psk}]
                                 [--ping-timeout-opt PING_TIMEOUT_OPT]
                                 [--show-ipsec-session-command SHOW_SA_CMD]
                                 [--ping-command PING_CMD] [-v] [-d] [--version]
    
    The IPsec end point health checker.
    
    optional arguments:
      -h, --help            show this help message and exit
      -s SERVER             specify the server address.
      -u USERNAME           specify the user name to execute the command.
      -p PASSWORD           specify the password for the user.
      --syslog              enable to send a message to syslog. default is
                            disable.
      --sshport SSH_PORT    specify the ssh port number. default is 22
      --sshmode SSH_MODE    specify the ssh authentication mode. Either 'pwd' or
                            'rsa' can be specified. default is 'rsa'
      --timeout TIMEOUT     specify the timeout to pint. default is 15
      --tagtype {rsa,psk}   specify the tag type to pick up the peer's IP address.
                            default is 'rsa'
      --ping-timeout-opt PING_TIMEOUT_OPT
                            specify the option name of the ping timeout. default
                            is '-W'
      --show-ipsec-session-command SHOW_SA_CMD
                            specify the command to show the ipsec session. default
                            is 'show crypto session'
      --ping-command PING_CMD
                            specify the ping command. default is '/sbin/ping'.
      -v                    enable verbose mode.
      -d                    increase debug mode.
      --version             show program's version number and exit
    
    ENVIRONMENT VARIABLE:
    
      Instead of specifying the username and password in the argument, the
      environment variable of HLCHK_SSH_USR and HLCHK_SSH_PWD can be used
      to specify the username and password respectively.
    
    NOTE:
    
      It does not look to define a general rule of the tag to pick the IP address
      of the peer assigned.  The known tags are like below:
    
      IKEv1 pre-shared key of ISR4000 IOS 16.5:
        IPSEC FLOW:.*host ([\w\d\.]+)
    
      IKEv2 RSA-SIG of ISR4000 IOS 16.5:
        Assigned address: ([\w\d\.]+)

## EXAMPLE

If there is no problem, it only shows the start and end time.

    % ipsec-health-check.py -s 10.71.158.190 -u admin -p 'lora,WAN-2017' --sshmode pwd --tagtype psk --ping-command=/usr/bin/ping 
    2017-09-25T07:53:42:ipsec-health-check:INFO: start ipsec-health-check
    2017-09-25T07:53:44:ipsec-health-check:INFO: end ipsec-health-check

If something happens on the route, you can see a warning message.

    % ipsec-health-check.py -s 10.71.158.190 -u admin -p 'lora,WAN-2017' --sshmode pwd --tagtype psk --ping-command=/usr/bin/ping 
    2017-09-25T07:57:22:ipsec-health-check:INFO: start ipsec-health-check
    2017-09-25T07:57:29:ipsec-health-check:WARNING: 192.168.1.36 is not stable.
    2017-09-25T07:57:29:ipsec-health-check:INFO: end ipsec-health-check

If you use the -v option, you can see a bit detail.

    % ipsec-health-check.py -s 10.71.158.190 -u admin -p 'lora,WAN-2017' --sshmode pwd --tagtype psk --ping-command=/usr/bin/ping -v
    2017-09-25T07:57:12:ipsec-health-check:INFO: start ipsec-health-check
    2017-09-25T07:57:19:ipsec-health-check:INFO: addr:192.168.1.36, tx:2, rx:0, loss:100.0%
    2017-09-25T07:57:19:ipsec-health-check:WARNING: 192.168.1.36 is not stable.
    2017-09-25T07:57:19:ipsec-health-check:INFO: end ipsec-health-check

You can defined the username and/or password in the environment variable
like below.

    % export HLCHK_SSH_USR="admin"
    % export HLCHK_SSH_PWD="password"

## TROUBLE SHOOTING

### not found in known_hosts

if you see the following line, you have to login into the server at once.

    paramiko.ssh_exception.SSHException: Server '10.71.158.190' not found in known_hosts

### No such file or directory

if you see the following line, you have to specify the valid ping command
by the --ping-command option.

    OSError: [Errno 2] No such file or directory

## Configuration example for rsyslog.

setup the rsyslog.conf to allow the UDP access.

    % sudo vi /etc/rsyslog.conf
    $ModLoad imudp
    $UDPServerRun 514

if you want rsyslog to only allow the access from the local system,
you have to define UDPServerAddress like below.

    % sudo vi /etc/rsyslog.conf
    $ModLoad imudp
    $UDPServerAddress 127.0.0.1
    $UDPServerRun 514

setup the config for this program to send the message into /var/log.

    % cat /etc/rsyslog.d/ipsec-health-check.conf 
    if $programname == 'ipsec-health-check' then /var/log/ipsec-health-check.log
    & ~

create the log file.

    touch /var/log/ipsec-health-check.log

Then, restart the rsyslog.

    service rsyslog restart

Enjoy.
