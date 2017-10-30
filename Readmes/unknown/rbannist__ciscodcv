## ciscodcv Dockerfile


This repository contains **Dockerfile** of ciscodcv for [Docker](https://www.docker.io/)


### Dependencies

* None


### Installation

1. Install [Docker](https://www.docker.io/) natively or [on MAC using boot2docker](https://docs.docker.com/installation/mac/) 

2. Download from private [Docker Hub](https://registry.hub.docker.com/u/rbannist/ciscodcv/): `docker pull rbannist/ciscodcv`

   (alternatively, you can build an image from Dockerfile: `docker build -t="rbannist/ciscodcv" github.com/rbannist/ciscodcv`)


### Usage

    docker run -i -t -p 8080:8080 rbannist/ciscodcv [bash]

### On a Mac running boot2docker

 1. VBoxManage controlvm boot2docker-vm natpf1 "web,tcp,,8080,,8080"

 2. boot2docker ssh ip addr show dev eth1 -> Get IP address

 3. Browse to http://'ip address':8080
