gandi-dyndns
----
This simple app, lets you dynamicaly update a DNS record on [Gandi](https://www.gandi.net) registrar using any [Cisco router](http://www.cisco.com/c/en/us/td/docs/ios-xml/ios/ipaddr_dns/configuration/15-mt/dns-15-mt-book/dns-dyn-dns-supp-ios.html#GUID-718C31D9-0D62-4675-B857-A20732374A12), Cisco ASA firewall (with some [restrictions](https://github.com/AlessioCasco/gandi-dyndns#cisco-asa-configuration)) or device that is able to send its pubblic IP and fqdn (hostname + domain name) through a GET or POST request like [dd-wrt](https://www.dd-wrt.com/wiki/index.php/Dynamic_DNS#Custom_.28URL_Updates.29), [openwrt](https://wiki.openwrt.org/doc/howto/ddns.client#custom_service,) [pfsense](https://doc.pfsense.org/index.php/Dynamic_DNS#Custom), [freenas](https://doc.freenas.org/9.3/freenas_services.html#dynamic-dns), [Synology](https://www.synology.com/en-uk/knowledgebase/DSM/help/DSM/AdminCenter/connection_ddns) and many others.

### From 0 to running

#### Download the app
* git clone https://github.com/AlessioCasco/gandi-dyndns.git

#### Installation
* Install bottlepy `pip install bottle`
* Rename or copy 'config-test.json' to 'config.json'.

#### Configuration

##### config.json
Default config looks like this:

```json
{"port":"8080",
"bind":"0.0.0.0",
"apikey":"your_gandi_apy_key",
"logging":{
    "log_enable":"false",
    "log_level":"INFO",
    "log_file":"./gandi-dyndns.log"}
}
```
- `port`- The HTTP port to listen on
- `bind`- The address that should be bound to for comunication. By default, this is "0.0.0.0", meaning gandi-dyndns will bind to all addresses on the local machine.
- `apikey`- Gandi apikey
- `log_enable`- Enable or Disable logging to file
- `log_level` - Log level to enable, possible values are: INFO, and DEBUG
- `log_file` - Log file relative or absolute path

##### gandi
In this example, we suppose you want to manage `router.example.com`
* You must have a zone file on you gandi account named as your domain. e.g. example.com needs a zone file called example.com, if you don't have it, create it and link it to your example.com domain [here](https://www.gandi.net/admin/domain/zone/list)

* Now edit the zone just created and add a new A record for the router subdomain:

|Field  | Value     |         |
| ------|:---------:|--------:|
| Type  | A         |         |
| TTL   | 5         | minutes |
| Name  | router    |         |
| Value | 127.0.0.1 |         |

* Once done, click on the button `use this version` to make the new zone file active.

#### Usage
##### Starting
Simply run the script

```bash
./gandi-dyndns.py
```

```bash
./gandi-dyndns.py -c configfile
```
This app accepts one optional parameter `-c, --config` that defines the location of the config file, by default this config file has to be in the same directory where `gandi-dyndns.py` is.

##### Interacion
Now your router, firewall or network appliance (for info about how to configure a cisco ASA firewall check the [config section](https://github.com/AlessioCasco/gandi-dyndns#cisco-asa-configuration)) can send updates to gandi-dyndns using `GET` or `POST` methods and the app will do the rest.

```
$machine_IP/DNS_name:$port/nic_update?ip=$IP&fqdn=$domain
```

To test the app manually (be aware that this may update your DNS name) issue this from your terminal:

```bash
curl -i "http:localhost:8080/nic_update?ip=1.1.1.1&fqdn=router.example.com
```
Or if you want to simulate a POST request:
```bash
curl -i -X POST "http:localhost:8080/nic_update?ip=1.1.1.1&fqdn=router.example.com
```
- `fqdn` This parameter is required and has to be the full FQDN of the device you want to update. e.g. router.example.com
- `ip`	 This parameter is optional and accepts only pubblic IP's. If none is supplied, the source address that generated the request is considered. This helps clients behind NAT or not able to send their IP to be used as well.


### HTTP status codes
* 200 => All good, 200 is given after updating the IP on Gandi and when there is no need to do so.
* 400 => Bad request, some parameters are missing, not formatted correctly or the provided IP is not a pubblic one.
* 404 => Not found, No domain found associated with the Gandi API, zone file missing or A record not found into the zone file.


### Monitoring
You can monitor if the app is up and running by simply send GET or POST requests to '/ping'

```bash
curl -i "http://localhost:8080/ping"

HTTP/1.0 200 OK
Date: Mon, 10 Apr 2017 22:05:08 GMT
Content-Length: 12
Server: gandi-dyndns
Content-Type: text/html; charset=UTF-8
Content-Type: text/html; charset=UTF-8

I'am alive!
```


### Cisco ASA configuration
Cisco ASA firewalls currently do not natively support DDNS Updates for HTTP-Based Protocols like routers do.
This simple hack is the only way I found to overcome this:

On your ASA firewall, under _config terminal_ mode, issue these two commands:
```
auto-update poll-period 30 5 1
auto-update server http://server:8080/nic_update?fqdn=router.example.com source outside
```
Note: Before entering the question mark (?) character, press the control (Ctrl) key and the v key together on your keyboard. This will allow you to enter the ? without the software interpreting the ? as a help query.

What your firewall will basically do is sending a POST request to your server that is running gandi-ddns every 30 minutes using the outside interface, if the request fails, it will try to send it again every minute for 5 times. Ip value is missing from the URL, so gandi will be updated with the source address of the request.

Details about the above commands can be found [here](http://www3.cisco.com/c/en/us/td/docs/security/security_management/cisco_security_manager/auto_update_server/4-9/user/guide/aus_ug_wrapper/as_boot.html#92739)


### Dependencies & Limitations

##### Dependencies
* [bottlepy](bottlepy.org)
* Gandi API. If you don't have it yet, enable the API from [Gandi](https://www.gandi.net/admin/api_key)

##### Limitations
* You must have a zone file on you gandi account named as your domain. e.g. example.com needs a zone file called example.com
* You can manage as many domains and subdomain as you want, but they all have to be owned by the same apikey.
* You will notice that gandi-dyndns sometimes needs quite a lot of time to respond with a 200 (~2s.), this is due to the slow nature of the Gandi API's.
* HTTPS is not available yet

##### Test
* Tested under Unix & Mac OS X using python 2.7.x
