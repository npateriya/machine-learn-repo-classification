# Overview

AcePruner is perl/MySQL-based tracking and reporting script for the Cisco ASA platform.

Its purpose is to track ACL hit counters over time.

# Installation

1.  Create a MySQL database

        /usr/local/mysql/bin/mysql -u root -p < acepruner.mysql
      
2.  edit acepruner.pl and configure the `sub init` subroutine, i.e. `vi acepruner.pl` 

#  Usage

1. Iteratively run `acepruner.pl`
2. Query the MySQL database to identify unused ACLs
      
        SELECT * FROM acepruner WHERE timestamp > (CURRENT_TIMESTAMP() - INTERVAL 5 MINUTE);
