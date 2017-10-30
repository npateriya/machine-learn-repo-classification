# Cisco APIC-EM REST API Learning Labs  
https://learninglabs.cisco.com/modules/net-controllers  
https://github.com/CiscoDevNet/apicem-1.3-LL-sample-codes  
https://github.com/CiscoDevNet/apicem-ga-1.2-ll-sample-code-basic-labs  
https://github.com/CiscoDevNet/apicem-ga-1.2-ll-sample-code-policy-labs.git  
https://github.com/CiscoDevNet/apicem-ga-ll-sample-code

## Table of contents

- [About](#about)
- [Demo](#demo)
- [Run from public docker image](#run-from-public-docker-image)
- [Build from scratch](#build-from-scratch)

## About  
Container that collects everything you need for the learning lab.  
Only requirement is to have docker on a terminal.  

No need to install different python version and no need to install python modules to do the lab

## Demo
![](https://raw.githubusercontent.com/robertcsapo/cisco-devnet-apicem-learning-labs/master/cisco-devnet-apicem-learning-labs.gif)

## Run from public docker image  

Without persistent storage  
```docker run --rm -it robertcsapo/cisco-devnet-apicem-learning-labs:latest```  

With persistent storage  
```docker run --rm -it -v "$PWD":/root/ robertcsapo/cisco-devnet-apicem-learning-labs:latest```  

## Got your own APIC-EM?  

Cool, then just use the environment variables for the docker container  
```
docker run --rm -it \
-e APICEM_IP=sandboxapic.cisco.com \
-e APICEM_USER=devnetuser \
-e APICEM_PASSWORD=Cisco123! \
robertcsapo/cisco-devnet-apicem-learning-labs:latest
```

## Build from scratch
```
git clone https://github.com/robertcsapo/cisco-devnet-apicem-learning-labs.git
cd cisco-devnet-apicem-learning-labs
docker build -t robertcsapo/cisco-devnet-apicem-learning-labs:latest .
```

## Start the Cisco APIC-EM REST API Lab  
https://learninglabs.cisco.com/modules/net-controllers

Lab Material in /root/
