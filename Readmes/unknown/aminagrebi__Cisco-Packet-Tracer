# Cisco Packet Tracer

![Image alt attribute](https://github.com/aminagrebi/Cisco-Packet-Tracer/blob/master/0002.PNG)

## Download
[Packet Tracer 32 bit](https://s3.amazonaws.com/lr_assets/20002/518210206/14350002/1.0?response-content-type=application/x-msdownload&response-content-disposition=attachment%3Bfilename%3D%22PacketTracer71_32bit_setup_signed.exe%22&AWSAccessKeyId=AKIAJ7XLQRRFSJBP4HZQ&Expires=1509180577&Signature=OmApeKTehv0%2BJl2tJ9aa8OUi11A%3D)	

[Packet Tracer 64 bit](https://s3.amazonaws.com/lr_assets/20002/518210206/14350004/1.0?response-content-type=application/x-msdownload&response-content-disposition=attachment%3Bfilename%3D%22PacketTracer71_64bit_setup_signed.exe%22&AWSAccessKeyId=AKIAJ7XLQRRFSJBP4HZQ&Expires=1509181351&Signature=i9pGwlQ32E8CUFrbgQCMK7cNq%2FA%3D)


## Router Configuration

1. Base Configuration to the Router
```
Router>enable

Router#configure terminal

Enter configuration commands, one per line. End with CNTL/Z.

Router(config)#hostname MassaRouter

MassaRouter(config)#ip domain-name massaelsham.com

MassaRouter(config)#banner motd # Authorized use only!

Enter TEXT message. End with the character '#'.

unauthorized access will be punished at the full extent of the law #

MassaRouter(config)#exit

MassaRouter#

%SYS-5-CONFIG_I: Configured from console by console

MassaRouter#exit

MassaRouter>

```

2. Securing Access to the Router
```
MassaRouter>enable

MassaRouter#configure terminal

MassaRouter(config)#line console 0

MassaRouter(config-line)#password cisco

MassaRouter(config-line)#login

MassaRouter(config-line)#line aux 0

MassaRouter(config-line)#password cisco

MassaRouter(config-line)#login

MassaRouter(config-line)#exit

MassaRouter(config)#exit
```
```
MassaRouter> enble

MassaRouter# config t

MassaRouter(config)#enable secret cisco

MassaRouter(config)#exit

MassaRouter#exit
```

```
MassaRouter> enble

MassaRouter# config t

MassaRouter(config)#show running-config

MassaRouter(config)#exit

MassaRouter#exit
```

```
MassaRouter> enble

MassaRouter# config t

MassaRouter(config)#service password-encryption

MassaRouter(config)#exit

MassaRouter#show run
```

```
MassaRouter>enable

MassaRouter#configure terminal

MassaRouter(config)#line console 0

MassaRouter(config-line)#logging sync

MassaRouter(config)#line aux 0

MassaRouter(config-line)#logging sync


MassaRouter(config-line)#exit

MassaRouter(config)#show running-config
```
3.Enabling SSH on the Router

```
MassaRouter> enble

MassaRouter# config t

MassaRouter(config)#crypto key generate rsa

The name for the keys will be: MassaRouter.massaelsham.com

Choose the size of the key modulus in the range of 360 to 2048 for your

  General Purpose Keys. Choosing a key modulus greater than 512 may take
  
  a few minutes.

How many bits in the modulus [512]: 1024

% Generating 1024 bit RSA keys, keys will be non-exportable...[OK]

MassaRouter(config)#ip ssh version 2

MassaRouter(config)#username amin secret cisco

MassaRouter(config)#line vty 0 4

MassaRouter(config-line)#transport input ssh

MassaRouter(config-line)#login local

MassaRouter(config-line)#logging sync

MassaRouter(config-line)#exit
```
4.Configuring interfaces

```
MassaRouter(config)#int f0/0
MassaRouter(config-if)#ip address 10.0.0.1 255.255.255.128
MassaRouter(config-if)#shutdown
MassaRouter(config-if)#no shutdown
MassaRouter(config-if)#int f1
MassaRouter(config-if)#ip address 10.0.0.129 255.255.255.128
MassaRouter(config-if)#no shutdown
```
5.Testing your Router
```
MassaRouter#copy running-config startup-config
```



