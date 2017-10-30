Cisco Command Line Interface
============================

A modular framework for implementing a Cisco-like CLI on a \*NIX system.
Arbitrary command trees can be defined using a simple syntax.

# Key Design Points

* Provides a modular CLI configuration using multiple tree definition files.
* Handles auto-completion of commands when pressing the TAB key.
* Provides context-sensitive help by pressing the ? key.
* Provides a secure interface to the environment. User will not get access to
  execute arbitrary commands on the system, but only those commands permitted by
  the privilege and group.
* Ability to be driven non-interactively.

# Implementation Mechanics

* Implemented as a dynamic library, allowing the system designer a method
  to dynamically build the interface programatically.
* The interface can also be built from a compiled description file, thereby
  obviating the need to recompile the executable every time the interface
  changes.

