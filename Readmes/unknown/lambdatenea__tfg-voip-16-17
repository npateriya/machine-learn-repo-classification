#  :mortar_board: Prototype IP telephony system based on free software

Code | https://github.com/lambdatenea/tfg-voip-16-17
---- | ----------------------------------------------
Last version | 1.0.4
Issues | https://github.com/lambdatenea/tfg-voip-16-17/issues/
Asterisk version | Asterisk 13.15.0
Cisco version | IOS 12.3 (c870)
Debian version | Debian GNU/Linux 8 (64 bit)

## What's this?

This is a final project of the University of Castilla la Mancha about the VoIP technology. 

## Description
In summary, the implementation of IP telephony systems in business environments brings many advantages, both economically and facilities. Therefore, in order to facilitate the migration of traditional telephony systems to IP voice systems, this final project proposes the development of a prototype that simulates a business environment with two remote locations, on which a system will be implemented Of low cost IP telephony based on free software. In this way, before testing the change of the system of a company, it will be possible to enhance tests of configurations, as well as demonstrations to those responsible. Figure 1.1 shows the prototype structure.

![Figure 1.1](https://github.com/lambdatenea/tfg-voip-16-17/blob/master/voip.png)

## Configuration

1. Install Asterisk in Debian:

Commands:

```
$ sudo apt-get install asterisk asterisk-core-sounds-es-gsm -y
```

Handly:

https://www.howtoforge.com/tutorial/how-to-install-asterisk-on-debian/

2. Install keepalived and fail2ban:
```
$ sudo apt-get install keepalived fail2ban
```

3. Download this project:
```
$ cd /usr/src
$ git clone https://github.com/lambdatenea/tfg-voip-16-17 o wget https://github.com/lambdatenea/tfg-voip-16-17
```

4. Copy files:
```
$ sudo cp /usr/src/asterisk/* /etc/asterisk
$ sudo cp /usr/src/debian/keepalived.conf /etc/keepalived/keepalived.conf
$ sudo cp /usr/src/debian/jail.conf /etc/fail2ban/jail.conf
$ sudo cp /usr/src/debian/copy_conference.bash /etc
```
5. Finally, setup the crontab:
```
$ crontab -e
```
Then it open a crontab file and you put this:
```
0 12 * * 1 bash /etc/copy_conference.bash
```

This task run every Monday at 12 o'clock.

## Licence

This project is licensed as BSD license.


## Author

Héctor Fernando Bahamonde Rivera.

## Bibliography
Asterisk. Contains description about the config's files and examples: https://wiki.asterisk.org/wiki/display/AST/Home
Voip-info.org. It's the same as the before: https://www.voip-info.org/wiki-Asterisk                                  
Asterisk Sounds. It's a tutorial to install the spanish's language: https://www.asterisksounds.org/es-es/instalar
