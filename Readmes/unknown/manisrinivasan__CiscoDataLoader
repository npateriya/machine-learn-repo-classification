CiscoDataLoader
===============

Create the keyspace and CF needed for this test program:

CREATE KEYSPACE cisco WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': '1'
};

USE cisco;

CREATE TABLE testtable (
  id int,
  created int,
  currency text,
  last_update int,
  points int,
  userid text,
  uuid text,
  version int,
  PRIMARY KEY ((id))
) 
  
  
Invoke Driver.java with the following Program arguments
servername  writes #partitions #clustercolsperpartition #batchsize #startpartition"
127.0.0.1 writes 100000  1 1 1
( this inserts 1m rows to a C* instance running on 127.0.0.1).


The logic to generate data is in Mysqldao.getCiscoOffers. - YOu may have to use your tables in Oracle and construct the List Collection 
Object.

The CiscoDaoModified.ciscoInsertData will then insert the data generated in the previous step into the C* Column Family.

		PreparedStatement ps = session
				.prepare("INSERT INTO cisco.testtable "
						+ "(id, uuid,version,created,last_update,currency,userid,points) "
						+ "VALUES (?,?,?,?,?,?,?,?);");
			.....// Use execute Async to load the data and bind the variables.
			results.add(session.executeAsync(
					ps.bind(cisco.getId(), cisco.getUuid(), cisco.getVersion(),
							cisco.getCreated(), cisco.getLast_update(),
							cisco.getCurrency(), cisco.getUserid(), cisco.getPoints())));	
						

This is a Maven project and the pom.xml has the right dependencies. Kindly change the version of the Datastax Driver to
( I had 2.0.0)

 <dependencies>
        <dependency>
            <groupId>com.datastax.cassandra</groupId>
            <artifactId>cassandra-driver-core</artifactId>
            <version>2.1.0</version>
        </dependency>
   
