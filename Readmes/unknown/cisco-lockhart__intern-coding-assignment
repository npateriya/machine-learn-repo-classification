One of Cisco Security's products is called [Advanced Malware Protection](http://www.cisco.com/c/en/us/products/security/advanced-malware-protection/index.html). 
At a very basic level, one of the things AMP does is to look at files users are 
downloading on to their corporate networks to see if the file is of a type that
needs to be either blocked, scanned, or allowed. Normally, this decision is made 
on the basis of:
* the policies configured by the network security administrator, and
* the threat intelligence Cisco holds within its servers.

As you can imagine, one of the challenges in implementing a system like AMP is the ability to determine
the action to perform upon receipt of a file. 

## Basic Requirements

In this exercise, we would like you to write a simple application that will determine
the action to perform when it is provided with a file. The application should:

* use Java 8 (bonus points for using the Streams API)
 
* run on the command line with one argument, the file name.

* preload a map of file MIME types and the action to perform for each of the MIME types when the application 
is started. Use this map to identify what to do with the file.We have provided 
a training JSON file with this assignment that you can use to do this.

* be built using Maven and produce a JAR file. The JAR file
should be executable.

The output of the application should be a single string that says:
  * allow
  * block
  * scan

If the file MIME type is unknown, you can default to the action `allow`.

## Additional requirements:

* Handle unexpected inputs (such as a missing file) gracefully.
* Log the output of the application using a logging system rather than to System.out
* Design the application such that we can replace the command-line application with, say,
a RESTful API. To wit, all of your business logic should **not** be in one file.

## Skeleton Project

To help you get started, we have provided you with a skeleton application that 
you can clone from

https://github.com/cisco-lockhart/intern-coding-assignment.git

To run all of the unit tests, run `mvn test`. When you submit your application,
I should be able to build the application using `mvn install`, which should produce
a JAR file I can run directly on my command line.

** NOTE **
There is a broken test here. Fix it.



## Libraries and tips

* To determine the MIME type of a file, you may want to look at Apache Tika, and
specifically the `detect(File file)` method.
* To parse the JSON we have provided, you may want to look at [Jackson](https://github.com/FasterXML/jackson-docs)
* If you need to learn more about Maven, check out this [tutorial](http://www.vogella.com/tutorials/ApacheMaven/article.html)
* As you will know from reading the tutorial, Maven uses dependencies to bring in
third-party tools. There is an example of how to do this in the skeleton project.
If you need to find a dependency, visit [search.maven.org](https://search.maven.org).
* To create an executable JAR using Maven, look at the Maven Assembly Plugin and perhaps search for how to build executable JARs
with dependencies.


Bonuses:
* Write unit tests for everything you build.

