# bosun

A Flask API to interact with Cisco iOS devices via Netmiko SSH that enables programatic configuration and management of Cisco iOS switches.

Big thanks to [Jason Edelman](https://github.com/jedelman8) and his [ntc templates](https://github.com/networktocode/ntc-templates)!

## HTTP API

### Endpoints

**show switch VLANs**<br>
GET `http://localhost/show-switch/vlan?host={{ host name }}`<br>

Response Body

```json
{
  "vlans": [
    {
      "name": "default",
      "status": "active",
      "vlan_id": "1"
    },
    {
      "name": "Infrastructure",
      "status": "active",
      "vlan_id": "10"
    },
    {
      "name": "trbrf-default",
      "status": "act/unsup",
      "vlan_id": "1005"
    }
  ]
}
```

**show switch interfaces status**<br>
GET `http://localhost/show-switch/interfaces-status?host={{ host name }}`<br>

Response Body

```json
{
  "interfaces": [
    {
      "duplex": "a-full",
      "name": "s2 node",
      "port": "Gi1/0/1",
      "speed": "a-100",
      "status": "connected",
      "type": "10/100/1000BaseTX",
      "vlan": "74"
    },
    {
      "duplex": "a-full",
      "name": "bench 01",
      "port": "Gi1/0/2",
      "speed": "a-1000",
      "status": "connected",
      "type": "10/100/1000BaseTX",
      "vlan": "trunk"
    },
    {
      "duplex": "auto",
      "name": "bench 02",
      "port": "Gi1/0/3",
      "speed": "auto",
      "status": "notconnect",
      "type": "10/100/1000BaseTX",
      "vlan": "76"
    }
  ]
}
```

**show google spreadsheet switch config**<br>
GET `http://localhost/show-config/google?spreadsheet_key={{ google spreadsheet key }}`<br>

Response Body

```json
{
  "config": [
    {
      "mode": "access",
      "name": "s2 node",
      "port": "gi1/0/1",
      "vlan": "74"
    },
    {
      "mode": "trunk",
      "name": "bench 01",
      "port": "gi1/0/2",
      "vlan": "76"
    }
  ]
}
```

**show github switch config**<br>
GET `http://localhost/show-config/github?git_user={{ github username }}&git_repo={{ github repo }}&git_branch={{ github branch }}&config_file={{ config file name with extension }}`<br>

Response Body

```json
{
  "config": [
    {
      "mode": "access",
      "name": "s2 node",
      "port": "gi1/0/1",
      "vlan": "74"
    },
    {
      "mode": "trunk",
      "name": "bench 01",
      "port": "gi1/0/2",
      "vlan": "76"
    }
  ]
}
```

**configure switch with google sheets spreadsheet template**<br>
GET `http://localhost/configure-switch/google?spreadsheet_key={{ google spreadsheet key }}?host={{ host name }}`<br>

Response Body

```json
{
  "message": "config term\r\nEnter configuration commands, one per line.  End with CNTL/Z.\r\nNY15m-3e(config)#interface gi1/0/1\r\nNY15m-3e(config-if)#description s8 node\r\nNY15m-3e(config-if)#switchport access vlan 74\r\nNY15m-3e(config-if)#switchport trunk encapsulation dot1q\r\nNY15m-3e(config-if)#switchport mode access\r\nNY15m-3e(config-if)#!\r\nNY15m-3e(config-if)#end\r\nNY15m-3e#"
}
```

**configure switch with github configuration file**<br>
GET `http://localhost/configure-switch/github?git_user={{ github username }}&git_repo={{ github repo }}&git_branch={{ github branch }}&config_file={{ config file name with extension }}?host={{ host name }}`<br>

Response Body

```json
{
  "message": "config term\r\nEnter configuration commands, one per line.  End with CNTL/Z.\r\nNY15m-3e(config)#interface gi1/0/1\r\nNY15m-3e(config-if)#description s8 node\r\nNY15m-3e(config-if)#switchport access vlan 74\r\nNY15m-3e(config-if)#switchport trunk encapsulation dot1q\r\nNY15m-3e(config-if)#switchport mode access\r\nNY15m-3e(config-if)#!\r\nNY15m-3e(config-if)#end\r\nNY15m-3e#"
}
```

**configure switch with json POST**<br>
POST `http://localhost/configure-switch/json?host={{ host name }}`<br>

Request Body

```json
[
	{
		"port": "gi1/0/1",
		"name": "port description",
		"vlan": "70",
		"mode": "access"
	},
	{
		"port": "gi1/0/2",
		"name": "port description",
		"vlan": "71",
		"mode": "trunk"
	}
]
```

Response Body

```json
{
  "message": "config term\r\nEnter configuration commands, one per line.  End with CNTL/Z.\r\nNY15m-3e(config)#interface gi1/0/1\r\nNY15m-3e(config-if)#description s8 node\r\nNY15m-3e(config-if)#switchport access vlan 74\r\nNY15m-3e(config-if)#switchport trunk encapsulation dot1q\r\nNY15m-3e(config-if)#switchport mode access\r\nNY15m-3e(config-if)#!\r\nNY15m-3e(config-if)#end\r\nNY15m-3e#"
}
```

**write memory**<br>
GET `http://localhost/configure-switch/write-mem?host={{ host name }}`<br>

Response Body

```json
{
  "message": "Building configuration...\n"
}
```
