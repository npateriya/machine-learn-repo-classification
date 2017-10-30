# CMAD Advanced Project

## A Question & Answers application (similar to StackOverflow) 

The minimally viable feature set include:

1. Ability to register and login
2. View a list of questions
3. Ask a question
4. View a specific question along with answer
5. Submit answers to questions
6. Ability to search for questions

### Application UI (SPA) built on React, Redux - 
http://35.184.123.212/#/

### Technology Stack
![Tech stack](/docs/screenshots/tech_stack.png)

### Architecture Diagram
![Arch diag](/docs/screenshots/arch_diag.png)

### GC VM Instances View
![gc_vm_instances diag](/docs/screenshots/gc_vm_instances.png)

### GC Build Triggers
![gc_build_triggers diag](/docs/screenshots/gc_build_triggers.png)

### GC Clusters
![gc_clusters diag](/docs/screenshots/gc_clusters.png)

### GC Load Balancers
![gc_loadbalancers diag](/docs/screenshots/gc_loadbalancers.png)

### REST API Details
https://app.swaggerhub.com/apis/narenderpal/cmad-app/1.2.0

### User Microservice 
https://github.com/narenderpal/user-service
User-Service provides the user registration, authentication and authorisation. Implemented using Vert.x Java. Uses MongoDB as persistence store. JWT Auth for generating access token for user authentication. 

### Question Microservice  
https://github.com/narenderpal/question-service
Question-service provides the functionality to view/post  questions/answers/comments. Implemented using Vert.x java. Used MongoDB as persistence store. 

### API Gateway Microservice 
https://github.com/narenderpal/api-gateway
API-Gateway is responsible for handling each request and based on the get/put/post header route it to auth handler and then dispatch it to each service endpoint (user or question service). Implemented using Vert.x java. 

### UI Service 
https://github.com/narenderpal/cmad-ui-app
UI-Service is the single page user interface application and the frontend door. Built on ReactJS , Redux, CSS, HTML, jQuery and using tools like Babel, Webpack.

### User Login UI
![ui_login diag](/docs/screenshots/ui_login.png)
### View Questions
![ui_questions diag](/docs/screenshots/ui_questions.png)
### View Answers
![ui_answers diag](/docs/screenshots/ui_answers.png)




## CMAD Staging Release Project
https://github.com/narenderpal/vertx-stackoverflow

### Build Pipeline
![Tech stack](/docs/screenshots/staging-build-pipeline.png)

### Docker Hub Repo
![Arch diag](/docs/screenshots/docker-hub-image.png)

### SonarQube Static Code Analysis
![gc_vm_instances diag](/docs/screenshots/sonar-qube-sca.png)



 
