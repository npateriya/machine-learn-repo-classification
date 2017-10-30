Take A Train
------------
[![Docker Repository on Quay Enterprise](https://containers.cisco.com/repository/paulvand/spvss-takeatrain/status "Docker Repository on Quay Enterprise")](https://containers.cisco.com/repository/paulvand/spvss-takeatrain)
[![Website](https://img.shields.io/website-up-down-green-red/http/shields.io.svg)](http://spvss-takeatrain.cisco.com/)

A Cisco teams "Train-o-scope" for the Infinite Home project.

Data are retrieved from the Rally project.



## HOWTO Install on an OpenStack VM

### Install Docker

```console
  sudo apt-get update
  sudo apt-get install apt-transport-https ca-certificates
  sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

  #sudo vi /etc/apt/sources.list.d/docker.list # remove all entries
  sudo echo "deb https://apt.dockerproject.org/repo ubuntu-trusty main" > /etc/apt/sources.list.d/docker.list
  sudo apt-get update
  #apt-cache policy docker-engine
  sudo apt-get install docker-engine
  sudo service docker start

  docker login containers.cisco.com

  sudo apt-get install git
```

### Clone this repository

### Build the image from Dockerfile

```console
  cd take-a-train/
  sudo docker build -t takeatrain .
```

### Start the image (exposing to port 80)

```console
  sudo docker run --rm -p 80:3000  \
    -e RALLY_API_KEY="_GeUV8XARSxa9zLefh..."  \
    -e MAILER_USER_PASSWORD="user:password"  \
    takeatrain
```
