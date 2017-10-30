# CPA Configurator


This simple Python script is an interactive configuration tool for Cisco IOS and IOS-XE devices aimed at providing a baseline IPsec configuration in line with the requirements of the UK CPA product evaluation scheme.

The script will generate either a Foundation or End-State (PRIME) VPN configuration.

# Configuration Coverage

The script will generate a configuration covering the following elements;

* IKEv1 or IKEv2 with associated crypto profile
* IPsec Profile and transform-set
* RSA or ECDSA keypair and certificate trustpoint
* Static VTI tunnel
* Configures a supplied enable password
* Disables Telnet and enables SSH version 2
* Disables numerous IOS services such as TCP Small Servers, CDP, IP Redirects etc.

The script can optionally take a certificate authority X.509 certificate as a PEM encoded file. This will be added to the script output to authenticate the configured trustpoint such that the user only has to then perform the 'enroll' step manually.

# Usage

The script can be run using `python configurator.py` and can optionally be passed a PEM encoded CA certificate using `python configurator.py ca_cert.pem`.

Once complete, the script will output the configuration wchih can then be copied and pasted in to the target device. Please make sure you're in 'enable' mode prior to pasting!

If no CA certificate is supplied then you will need to manually run `crypto pki authenticate` and `crypto pki enrol` against the trustpoint (which the script will name `cpa_ca`)

# To Do

This is a basic script and there are a number of things I'd like to improve;

* Make the script interact directly with an IOS device to push the config directly
* Allow advanced customisation such as certificate subject-name, use of already configured trustpoints etc.
* Support ASA VPN configuration
