cisco_vpfa
=======

#### Table of Contents

1. [Overview - What is the cisco_vpfa module?](#overview)
2. [Module Description - What does the module do?](#module-description)
3. [Setup - The basics of getting started with cisco_vpfa](#setup)
4. [Implementation - An under-the-hood peek at what the module is doing](#implementation)
5. [Limitations - OS compatibility, etc.](#limitations)
6. [Development - Guide for contributing to the module](#development)
7. [Contributors - Those with commits](#contributors)

Overview
--------

The cisco_vpfa module is a part of [OpenStack](https://www.openstack.org), an effort by the OpenStack infrastructure team to provide continuous integration testing and code review for OpenStack and OpenStack community projects not part of the core software.  The module its self is used to flexibly configure and manage the FIXME service for OpenStack.

Module Description
------------------

The cisco_vpfa module is a thorough attempt to make Puppet capable of managing the entirety of cisco_vpfa.  This includes manifests to provision region specific endpoint and database connections.  Types are shipped as part of the cisco_vpfa module to assist in manipulation of configuration files.

Setup
-----

**What the cisco_vpfa module affects**

* [Cisco_vpfa](https://wiki.openstack.org/wiki/Cisco_vpfa), the FIXME service for OpenStack.

### Installing cisco_vpfa

    cisco_vpfa is not currently in Puppet Forge, but is anticipated to be added soon.  Once that happens, you'll be able to install cisco_vpfa with:
    puppet module install openstack/cisco_vpfa

### Beginning with cisco_vpfa

To utilize the cisco_vpfa module's functionality you will need to declare multiple resources.

Implementation
--------------

### cisco_vpfa

cisco_vpfa is a combination of Puppet manifest and ruby code to delivery configuration and extra functionality through types and providers.

Limitations
------------

* All the cisco_vpfa types use the CLI tools and so need to be ran on the cisco_vpfa node.

Beaker-Rspec
------------

This module has beaker-rspec tests

To run the tests on the default vagrant node:

```shell
bundle install
bundle exec rake acceptance
```

For more information on writing and running beaker-rspec tests visit the documentation:

* https://github.com/puppetlabs/beaker-rspec/blob/master/README.md

Development
-----------

Developer documentation for the entire puppet-openstack project.

* https://docs.openstack.org/puppet-openstack-guide/latest/

Contributors
------------

* https://github.com/openstack/puppet-cisco_vpfa/graphs/contributors
