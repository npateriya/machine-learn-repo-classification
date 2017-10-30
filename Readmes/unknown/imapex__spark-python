# spark-python

[![Build Status](http://drone.green.browndogtech.com/api/badges/imapex/spark-python/status.svg)](http://drone.green.browndogtech.com/imapex/spark-python)

# Description

spark-python is a set of libraries for working with the Cisco Spark API's.

# Installation

## Environment

Prerequisites

* Python 2.7+
* [setuptools package](https://pypi.python.org/pypi/setuptools)

## Option A - pip

    pip install spark-python

## Option B - from source:

If you have git installed, clone the repository

    git clone https://github.com/imapex/spark-python
    cd spark-python
    python setup.py install


Provide instructions on how to install / use the application

# Usage

You can find samples in the [./samples](./samples) directory.

The basic object you will use to interact with the API is an instance of Session

    from spark.session import Session
    session = Session('https://api.ciscospark.com', 'YOURTOKENHERE')

To see a list of rooms you are a member:

    from spark.rooms import Room
    room = Room.get(session


# Development

This is an early in life project that welcomes community contributions, if you feel there is a use case
the library doesn't cover, go ahead and send us a PR!!  A few things to keep in mind:



## Linting

We use flake 8 to lint our code. Please keep the repository clean by running:

    flake8

## Testing

The IMAPEX team attempts to have unittests with  100% code coverage. The main test suite is located in [./tests/testing.py](./tests/testing.py)
You will need to provide an API token

The tests are can be run with the following command

    python testing.py


When adding additional code or making changes to the project, please ensure that unit tests are added to cover the
new functionality and that the entire test suite is run against the project before submitting the code.
Minimal code coverage can be verified using tools such as coverage.py.

For instance, after installing coverage.py, the toolkit can be run with the command::

    coverage run tests.py

and an HTML report of the code coverage can be generated with the command::

    coverage html


# License


TBD
