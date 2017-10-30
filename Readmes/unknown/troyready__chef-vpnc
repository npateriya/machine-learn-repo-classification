# vpnc

Deploys Maurice Massar's 'client for cisco vpn concentrator'. Primarly designed to provide a temporary VPN connection for remote infrastructure testing.

## Usage

Set the attributes below and include `vpnc::default` at the top of your node's run_list.

## Attributes

- `node['vpnc']['ipsec_gateway']` - value for "IPSec gateway"
- `node['vpnc']['ipsec_id']` - value for "IPSec ID"
- `node['vpnc']['ipsec_secret']` - value for "IPSec secret"
- `node['vpnc']['xauth_user']` - value for "Xauth username"
- `node['vpnc']['xauth_pass']` - value for "Xauth password"

Note: the `ipsec_secret` and `xauth_pass` values may be obfuscated via base64 encoding; simply prefix their encoded values with the fixed string `base64:`. E.g.:

```
<start irb>
require 'base64'
puts "base64:#{Base64.encode64('imsosecret5').chomp}"
# base64:aW1zb3NlY3JldHk1
```

- `node['vpnc']['compile_time']` - Should vpnc be deployed at compile time? Defaults to `false`
