# parse-slbl
Utility to parse safe/blocklist backup data from Cisco IronPort into various reporting formats.

## Motivation
I needed a quick tool to take the safe/block list export from a Cisco Email Security (IronPort) appliance and pivot the data into something I could manipulate with Excel or other tools. The backup file is a simple CSV file, but does not lend itself well as each row is a variable length record containing a user's safe or block list.

This utility takes that data and creates two different CSV files that display the safe/block lists in a more usable format for reporting.

## Installation

### Requirements
  * Perl 5.x

### Setup and Configuration
There are no special steps except to put the script on your computer and execute it with a Perl interpreter.

## Usage

```sh
perl parse-slbl.pl <filename>
```
For ``<filename>``, specify the csv file from your Cisco IronPort device.

## License
**parse-slbl** is licensed under the [MIT license.](https://github.com/seliger/parse-slbl/blob/master/LICENSE)
