# Cisco RMA Console

# Description

This is a quick proof of value application demonstrating the use of Cisco's API
for integration with the RMA system. It requires a Cisco issued OAuth2 keypair
in order to be installed and used.

This repository includes scripts for installation into an instance of [mantl](http://mantl.io).

# Installation

## Environment

Prerequisites

* Python 2.7+
* [setuptools package](https://pypi.python.org/pypi/setuptools)
* [Flask](http://flask.pocoo.org)
* [Requests](http://docs.python-requests.org/en/master/)
* Hosting environment - any environment capable of hosting multiple containers
and implementing networking between them. Tested/supported environments are
currently [mantl](http://mantl.io) and [Docker](http://www.docker.com).

The front end application makes use of [JQuery](http://jquery.com);
separate installation of these libraries is not required as they are linked from
public CDN networks.

## Downloading

Option A:

If you have git installed, clone the repository into a local directory.

    git clone https://github.com/sluzynsk/rmaconsole

Option B:


The latest build of this project is also available as a Docker image from Docker Hub

    docker pull sluzynsk/rmaconsole:latest




## Installing

### Mantl installation

These instructions assume that you have cloned the installation repo as listed
above.

* Edit the sample-rmaconsole.json file (included) to reflect credentials needed.
* Run the app_install.sh script to install the application to your
[mantl](http://mantl.io) server. This script will also create json files that can
be used to re-install the application if needed.

### Docker installation

Installation into a swarm should work but has not yet been tested.

# Usage

Simply enter one or more comma separated RMA numbers into the text box.


## Linting

We use flake 8 to lint our code. Please keep the repository clean by running:

    flake8

## Testing

Currently test coverage is lacking for this application.


# License

MIT license. See the included [LICENSE.TXT](LICENSE.TXT) file for details.
