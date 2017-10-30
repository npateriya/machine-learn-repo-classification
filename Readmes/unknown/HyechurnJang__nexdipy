# nexdipy
Cisco Open NX-OS Python API

**NEX**us-os **D**eveloping **I**nterface for **PY**thon

version : 0.1.1

last change : add model & pod & multi-pod & CLI

![Relations](./doc/Relation.png)

## Support Object

**Prepared Objects**

| NXOS Object Name | Nexdipy Object Name | Reserved Code Name | Description |
|------------------|---------------------|--------------------|-------------|
| topSystem | nexSystemObject | System | System Description |
| interfaceEntity | nexInterfaceObject | Interface | Interface Entity |
| l1PhysIf | nexPhysIfObject | PhysIf | Physical Interfaces |
| pcAggrIf | nexAggrIfObject | AggrIf | Aggregated Interfaces |
| l3Inst | nexContextObject | Context | L3 Context (VRF) |

**And Retrieve Anything with Node Object through "Class()" Method**

## Install

**From GIT**

	$ python setup.py build
	$ python setup.py install

**From PIP**

	$ pip install nexdipy

## Example Acidipy

	import nxosdipy
	
	node = nxosdipy.Node('xxx.xxx.xxx.xxx', 'user', 'password') # Get node connection
	
	node.System.Interface.PhysIf.list() # Retrive physical interface list
	
	node.close() # Close node connection
