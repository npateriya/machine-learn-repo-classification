## The repository is branched off from:

https://sanrath@bitbucket.org/sanrath/mediatek_mt7610u_sta_driver_linux-64bit.git

## How to use this driver ##

* Download the sources from master branch 
MediaTek_mt7610u_STA_driver_Linux 64bit.zip

* Unzip and cd MediaTek_mt7610u_STA_driver_Linux 64bit
* run the following commands

```
#!shell

make clean
sudo make 
sudo make install 

```

Connect Linksys AE6000 Wifi Dongle to your PC / Laptop.. 

For other chipsets some files need modification including the following

* Makefile
* os/linux/config.mk 
* common/rtusb_dev_id.c

## Additional startup script is added, which may be needed ##

Add wireless_setup.sh to the system startup scripts (eg: /etc/init.d/wireless_setup.sh and link it with "ln -s /etc/init.d/wireless_setup.sh /etc/rc5/S99wireless_setup.sh" )
