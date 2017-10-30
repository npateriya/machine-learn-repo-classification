Overview
========

Allows to backup the Cisco devices



Config of the TFTP server
=========================

 apt-get install tftpd -y

 cat >/etc/xinetd.d/tftp << END

 service tftp

 {

 protocol        = udp

 port            = 69

 socket_type     = dgram

 wait            = yes

 user            = nobody

 server          = /usr/sbin/in.tftpd

 server_args     = /tftpboot

 disable         = no

 }
 
 END


To add device backups
=====================

For each device to add:

 - todo on the cisco router:
      - en
      - conf t
      - snmp-server community 'rw_community' RW
      - snmp-server host 'TFTPD_IP_address' version 2c 'rw_community'  config
      - exit
      - wr

 - todo on the hp switch:
      - conf t
      - snmp-server community 'rw_community' unrestricted
      - wr mem
      - exit

 - todo on the tftp server: >/tftpboot/{cisco,hp}.{'device_ip_addr','device_domain_name'}.conf

 - and, optionally: cisco.backup.sh -i >>/etc/iptables.d/cisco.backups.fw



That's all folks!

The new device will be automatically supported ...


USE
===

   # help:

   cisco.backup.sh -h

eg.

   # setup the list of cisco devices to backup:

   cisco.backup.sh --setupfrom /media/cloudfuse/os-backup/cisco.conf/

 or:

   touch /tftpboot/cisco.192.168.0.1.conf /tftpboot/hp.devname.my.domain.conf

   # cron example:

   00 00 * * *	ROOT=/media/cloudfuse/os-backup; curl -sfL https://github.com/peychart/snmp-cisco-backup/raw/master/cisco.backup.sh| /bin/bash -s -- -c community_name -d /tftpboot| while read i; do cp -f /tftpboot/$i $ROOT/cisco.conf/$i.$(date "+\%Y\%m\%d.\%H\%M") 2>/dev/null; done; cd /tftpboot/ && [ -d .git ] && (git add *.conf && git commit -m "New backups" && git push) >/dev/null 2>&1

