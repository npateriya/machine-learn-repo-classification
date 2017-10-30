# nxuti
nxuti is a CLI utility for talking to the Cisco Nexus NXAPI for managing Nexus Switch.

This CLI utility is capable of managing trunk and native vlans beneath ethernet
and port-channel interfaces.   VLAN objects can also be managed.

clone the project: $ git clone https://github.com/caboucha/nxuti.git

cd into the project dir. $ cd nxuti/cmd/nxctl

Pull in deps: $ go get github.com/spf13/cobra

Build the bin: $ go build -i

Run the client against the nx-api, passing the appropriate flags: $ ./nxctl trunkVlan show --interface ethernet -hosts=10.1.1.254 -user=admin -pass=cisco

Sample Usage below:

go run cmd/nxctl/nxutil.go trunkVlan show --interface ethernet --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk   Mode    State   Descr
eth1/33         vlan-1  1-4094  access  up
eth1/34         vlan-1  1-4094  access  up
<SNIP>

go run cmd/nxctl/nxutil.go trunkVlan show --interface ethernet:1/19 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk   Mode    State   Descr
eth1/19         vlan-1  None    trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan add --interface ethernet:1/19 --nativeVlan 200 --trunkVlan 200-203,206 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native          Trunk           Mode    State   Descr
eth1/19         vlan-200        200-203,206     trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan remove --interface ethernet:1/19 --nativeVlan 200 --trunkVlan 200-203,206 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk   Mode    State   Descr
eth1/19         vlan-1  None    trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan replace --interface ethernet:1/19 --trunkVlan 200-203,205 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk           Mode    State   Descr
eth1/19         vlan-1  200-203,205     trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan replace --interface ethernet:1/19 --trunkVlan 193-194 --nativeVlan 193 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native          Trunk   Mode    State   Descr
eth1/19         vlan-193        193-194 trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan add --interface ethernet:1/19 --nativeVlan 200 --trunkVlan 200-203,206 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native          Trunk                   Mode    State   Descr
eth1/19         vlan-200        193-194,200-203,206     trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan remove --interface ethernet:1/19 --nativeVlan 200 --trunkVlan 200-203,206 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk   Mode    State   Descr
eth1/19         vlan-1  193-194 trunk   up      connection to UCS bxb-ds-46

go run cmd/nxctl/nxutil.go trunkVlan replace --interface ethernet:1/19 --nativeVlan None --trunkVlan None --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
ID              Native  Trunk   Mode    State   Descr
eth1/19         vlan-1  None    trunk   up      connection to UCS bxb-ds-46

2) Do trunk/native vlan show/add/replace/remove operations on interface port-channel 55
Same as ethernet excep you use --interface port-channel:55 for example

3) Do Vlan show/add/remove operations
go run cmd/nxctl/nxutil.go vlan show --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI                     Admin State     OperState       Name
        193             unknown                 active          down            VLAN0193
        505             unknown                 active          down            VLAN0505
        101             unknown                 active          down            VLAN0101

go run cmd/nxctl/nxutil.go vlan show --vlanid 222 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI                     Admin State     OperState       Name

go run cmd/nxctl/nxutil.go vlan add --vlanid 222 --vnSegment 3333 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI                     Admin State     OperState       Name
        222             vxlan-3333              active          down            VLAN0222

go run cmd/nxctl/nxutil.go vlan remove --vlanid "222" --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI             Admin State     OperState       Name

go run cmd/nxctl/nxutil.go vlan add --vlanid 222 --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI             Admin State     OperState       Name
        222             unknown         active          down            VLAN0222

go run cmd/nxctl/nxutil.go vlan remove --vlanid "222" --hosts your-nexus-ip-address --user "admin" --pass "yourpassword"
        ID              VNI             Admin State     OperState       Name


