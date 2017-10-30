Using Cisco Shell you can connect to a remote Cisco device and execute a set of configurational commands.

It is a script I wrote for myself while working as a networks engineer in a local bank.

Installation
=============

You need CoffeeScript installed to run this script

```
npm install --global coffee
```

Also you need the "ssh2" module installed

```
npm install ssh2
```

And "sync" module installed (along with fibers)

```
npm install sync
```

Run it like this
=============

```
coffee run.coffee "{ \"user\": \"Username\", \"password\": \"P@$$w0rD\", \"device\": \"4.4.4.4\" }"
```

Syntax
=============

The executed commands are located in the "script.coffee" file.

This is a usual "CoffeeScript" script, nothing complicated.

You can use functions, 3rd party libraries - anything - there.

$ is to execute a single command.

$$ is to execute a batch of commands sequentially.

These two functions can also accept "callbacks" (with the output of the commands).

Output example
=============

```
Connection :: connect
Connection :: ready
===================================================================
> show ip interface GigabitEthernet 0/0
===================================================================
GigabitEthernet0/0 is administratively down, line protocol is down
  Internet protocol processing disabled

Gigabit Ethernet 0/0 interface found. Proceeding.
Gigabit Ethernet 0/0 interface is switched off.
Here we could easily execute a command to bring it up if we wanted.
===================================================================
> show inventory
===================================================================
NAME: "CISCO2901/K9", DESCR: "CISCO2901/K9 chassis, Hw Serial#: FCZ1641617Y, Hw Revision:
1.0"
PID: CISCO2901/K9      , VID: V05 , SN: FCZ1641617Y

NAME: "4 Port FE Switch on Slot 0 SubSlot 0", DESCR: "4 Port FE Switch"
PID: HWIC-4ESW         , VID: V01 , SN: FOC14144F4E

NAME: "C1941/C2901 AC Power Supply", DESCR: "C1941/C2901 AC Power Supply"
PID: PWR-1941-2901-AC  , VID:    , SN:
===================================================================
> configure terminal
===================================================================
Enter configuration commands, one per line.  End with CNTL/Z.
===================================================================
> voice service voip
===================================================================
> allow-connections h323 to h323
===================================================================
> allow-connections h323 to sip
===================================================================
> allow-connections sip to h323
===================================================================
> supplementary-service h450.12
===================================================================
> redirect ip2ip
===================================================================
Connection :: end
Connection :: close
```

Known issues
=============

Sometimes (at least on my machine) it's stuck on "Connection :: connect" and then timeouts. I don't know why.