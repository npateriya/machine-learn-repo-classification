This is a simple implementation of the Standup Bot for Cisco Spark and HipChat Server.

It can function behind a firewall, so it is particularly useful for HipChat server installations that can not be accessed from the public internet.
 
The bot provides the same functionality between Cisco Spark and Hipchat, which could be helpful for teams wishing to maintain the same standup functionality as they migrate between platforms.
 
Be careful that you create a separate user to function as the Standup Bot and for testing. The integration test will wipe all rooms for the test user.

A walkthrough video of the implementation of the Bot can be found at: https://youtu.be/OscaWUFJcH4

#Getting started

The following environment variables need to be configured for both running integration tests and production operation:

````
ciscoSparkApiKey=YOURKEY
ciscoSparkApiKeyEmailAddress=test@example.org
ciscoSparkWebhookTargetUrl=http://example.org/test
````

##Building
````
mvn package -DciscoSparkApiKey=YOUR_TEST_USER_KEY -DciscoSparkWebhookTargetUrl=http://example.test.org/cisco-spark-standup -DciscoSparkApiKeyEmailAddress=standupbot-test@example.org
````

##Running
````
java -jar -DciscoSparkApiKey=YOUR_KEY -DciscoSparkWebhookTargetUrl=http://example.org/cisco-spark-standup -DciscoSparkApiKeyEmailAddress=standupbot@example.org target/standup-0.0.1-SNAPSHOT.jar 
````



