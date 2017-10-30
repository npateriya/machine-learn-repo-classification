# wsma_python
This Python module shows some examples of interacting with the WSMA (Web Services Management Agent) protocol on Cisco devices.

WSMA is supported on most Cisco enterprise devices (routers/switches), and first shipped in 2007. WSMA can run over HTTP(S) or SSH.

## Installation

```
# a virtualenv installation
$ virtualenv env
$ source env/bin/activate

# now copy the code from GitHub
$ git clone https://github.com/aradford123/wsma_python
$ cd wsma_python

# install it in the virtual environement
$ python setup.py install

# can also be done in one step:
$ pip install git+https://github.com/aradford123/wsma_python.git

# update examples/wsma_config.py to point to your device, with appropriate credentials

$ ./examples/00show_run.py
```

You can also use the script `./enable_wsma.py <ip> <username> <password>` to configure WSMA if required on your device. This configures WSMA over HTTPS, but you can also use HTTP or SSH as the transport.

You can also just paste in the following config snippet. This configures the HTTPS transport (devices also support SSH or HTTP) and local authentication (device can also support AAA).

```
ip http secure-server
ip http authentication local
wsma agent exec
 profile WSMA
wsma agent config
 profile WSMA
wsma profile listener WSMA
 transport https
end
```

## Examples
Contains a number of examples uses of the module.  
