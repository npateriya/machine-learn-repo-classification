
# SSH on VPN

## Usage

Must use `--cap-add=NET_ADMIN --device=/dev/net/tun` or `privileged` option.

Require: `VPN_SERVER` and `SSH_SERVER` environment variables.

```sh
$ docker run -it -e VPN_SERVER=vpn.exemple.com -e SSH_SERVER=ssh_user@ssh.example.com --privileged ornew/vpn-ssh
```

### Example

In `./volume/.ssh/config`:

```
# Define host
Host example
    HostName 0.0.0.0
    Port 22
    User ssh_user
    IdentityFile ~/.ssh/id_rsa
```

In `./volume/.openconnect/config`:

```
# VPN options
# see: http://www.infradead.org/openconnect/manual.html
user=vpn_user
```

In `./.env`:

```
VPN_SERVER=vpn.example.server.com

# If defined host at ~/.ssh/config
SSH_SERVER=example
# else: ssh option
#SSH_SERVER="ssh_user@ssh.example.com -p 22"
```

On `./`:

```sh
$ docker run -it --env-file=.env -v=$(pwd)/volume:/root --privileged ornew/vpn-ssh
```
