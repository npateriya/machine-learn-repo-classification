pix_to_pan
==========

Rudimentary ruby script to convert cisco pix ACLs to xml fragments for Palo Alto firewalls

DON'T USE AS-IS

This has been tested converting ACLs from an ASA5500 to a PA-5060. That's it.

It makes lots of assumptions:

* that you want your acls grouped
* that you have two zones
* that your "inside" zone can be identified by a single ip/netmask
* it assumes you are on a unix machine
* ...

It has lots of bugs:

* may incorrectly reorder your ACLs
* may incorrectly group your ACLs
* the human readable rule names it generated were too long for PA and replaced by opaque strings
* ...

It's provided because it may ba a useful starting point.

To run:

1. clone it
2. gem install treetop
3. ./pix_to_pan.rb ACL_File1 ..

It produces (on standard out) three sets of XML fragments for import on a PA: Adress Objects, Service Objects, and Rules.
