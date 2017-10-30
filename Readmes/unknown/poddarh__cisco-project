## Records API 

#### Overview 
This is an API to persist and retrieve arbitrary data. 

##### Dependencies Required 
- Java 8 
- MongoDB 2.6 or higher 

##### Build Instructions 
- __Linux/Mac__ 
  - Extract the `records.zip` file 
  - Build the project by running `./build` 
 

- __Windows__ 
  - Extract the `records.zip` file 
  - Build the project by running `build.cmd` 
 
##### Instructions to Run 
- __Linux/Mac__ 
  - Upon successful build, run `./start` *(Note: you may need to run as super user as it uses port `80`. Alternatively, you can modify the `start` file to run on another port.)* 


- __Windows__ 
  - Upon successful build, run `start.cmd` to start the server 
 
#### Improvements and Features 
- Error messages should have a `statusCode` field which will be known to both the server and the client and so it will be easier for the client to comprehend as to what went wrong. This will also enable us to externalize the error messages. 
- Depending on the way the application is going to be used, it may be useful to add security to restrict unauthorized usage. `uid` sounds to me like an acronym for user identifier and I would like to restrict users from editing other records. 
- I would also enable Cross-Origin Resource Sharing (CORS) if the client web application is going to be on a separate domain. 
- Again depending on the use case of the application, it may be useful to define set of fields it may accept instead of accepting arbitrary fields as it will make it easier to analyze or process the data. 
- I would use the HATEOAS constraint of the REST architecture to even more loosely couple the client and the server. The client will then only interact with the server through hypermedia provided by the server. This increases the ability to modify code on the server without worrying about the client breaking. 
- This may be totally out of context, but one use of storing arbitrary data I can think of is to be searched later. I would enable the ability to look for records without the uid but also by searching within the fields. To enable this, I would use a search engine library such as Apache Lucene. 
 