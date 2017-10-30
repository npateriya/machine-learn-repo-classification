# CiscoRESTAPI
An application project for Cisco.

It's just a simple RESTful API that deals with arbitrary JSON obejcts.
I used Python, the cherrypy http server library, an SQLite3 database,
git for version control, and Amazon aws for tier services.

To get started:
    1) Run check_modules.py to check if you're missing any python
    dependencies. If you are, install them.

    2) Go into cherrypy.conf and set sockey_host and socket_port to
    the configuration you want for your machine.
        - Host should be your public DNS or localhost
        - Port should be 80 because that's the usual for HTTP

    3) Run run.py [-d database_name]
        - You'll porbably need sudo if you're using your public IP
        - Don't worry about creating log or database files,
          these are created automatically. (You can enter a new
          database name and it will create the database for you)
        - I'd suggest running it in its own terminal, though in the
          background would work to.
    
    4) While run.py is going, run unit_tests1.py <your url> to test 
    basic functionality
        - YOU MUST GO INTO unit_tests1.py and near the top, manually
          change 'url' to be your url. You may need to add the port to
          the end.
        - I recommend doing this for both a database and non-database
          version of "run.py".
