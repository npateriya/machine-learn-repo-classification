# aci-starwars-demo
ACI programmability demo for Cisco Live 2016 Melbourne, Devnet. The force is strong with this demo.

# Environment

* Download & install [VirtualBox](https://www.virtualbox.org/) and [Vagrant](https://www.vagrantup.com/downloads.html)

# Installation
```
  git clone https://github.com/miciesla/aci-starwars-demo.git
  cd aci-starwars-demo
  vagrant up
```
# Usage
```
  vagrant ssh
  cd /vagrant
```
* Create the BB-9 tenant:
```
  python devnet-create-bb9.py
```
* Web browse to the load balancer address. Hit refresh to change the Millennium Falcon's charter.

* Create ACI tenancies for the other Star Wars characters
```
  python devnet-loop.py
```

* Remove all tenancies:
```
  python devnet-delete.py
```
