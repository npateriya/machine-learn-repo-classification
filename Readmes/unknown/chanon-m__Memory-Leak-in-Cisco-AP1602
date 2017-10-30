# Memory-Leak-in-Cisco-AP1602

In case you have a memory leak issue in Cisco AP1602, you can you this script to detect memory leak in Cisco AP1602 and reload IOS to have free memory. 

#Licensing Information: READ LICENSE

#Project source can be downloaded from
##https://github.com/chanon-m/Memory-Leak-in-Cisco-AP1602.git

#Author & Contributor

Chanon Mingsuwan

Reported bugs can be sent to chanonm@live.com

#How to run file

* Download files to your server

```

# git clone https://github.com/chanon-m/Memory-Leak-in-Cisco-AP1602.git

```

* Copy memleak.pl to /etc

```

# cp ./Memory-Leak-in-Cisco-AP1602-/memleak.pl /etc

```

* Make a file executable  

```

# chmod 755 /etc/memleak.pl

```

* Create a crontab job on your server every two hours

```

# crontab -e

0 */2 * * *      /etc/memleak.pl >> /var/log/memleak.log&

```
