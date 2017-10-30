# spark-space-monitor
Spark Space Monitor enhances Cisco Spark by notifying users when people ENTER and EXIT Spaces

When you need to know who ENTERED or EXITED a Cisco Spark Space that is part of a team as of today, the only way to do so is to proactively got to the Space and check PEOPLE.

This app fixes that. 

Any user within the Space can request to be notified whenver someone ENTERS or EXITS the Space by sending a message to the app with the /addme modifier

When one of these things happens (ENTERS or EXITS Space) the the app will do 2 things:
  1. Send a 1:1 Message to each person that requested to be notified.
  2. Send a message to the Space so everyone knows about this.
  
I added other options, some to enhance this basic functionality, and others to enhance the Space itself.

I will document these in details in a few days.

The one thing that the app missess right now is the CREATION of the TABLE in mySql if it is not there. It assumes it is there.

Until I create the TABLE within the app any developer will have to create the TABLE with this command (the name of the TABLE CANNOT be modified):

      CREATE TABLE SparkSpaceMonitor (id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, 
      personEmail VARCHAR(50),
      personId VARCHAR (100),
      personDisplayName VARCHAR (100),
      roomId VARCHAR(100),
      roomTitle VARCHAR(50),
      avatar VARCHAR(10),
      creation_date VARCHAR(40),
      modification_date VARCHAR(40));

