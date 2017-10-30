# Angular 1.5 Course

This repository holds the files from the course about Angular 1.5.
The course started at December 2016 and the lecture was Yariv Katz. 
The lesson material is located at [nerdeez.com](https://www.nerdeez.com)

## Day 1

#### History

- Internet History
- The Birth of Ajax Technology
- What is JQuery
- What is SPA
- The challenges of SPA
- The birth of frontend frameworks and Angular1.5
- What is the shortcoming of using a framework
- Angular2

#### Angular Architecture - Extended Hello World

The files in this topic demo are located at the folder angular-architecture-hello-world
In this demo we will create the following application:
* A single module that holds the entire application and bootstrapped to the DOM
* A template with a form and a text input to add a greeting to the list. 
Below the form there is a list of greeting message that the user can add a 
greeting by submitting the form
* A controller which grabs the list of greeting and displays it to the user
and also handles the form submission.
* A Factory which handles the data of our app which is an array of greetings

We will cover the following topics: 
- npm
- bower
- Angular modules
- Angular templates introduction
- Angular controller introduction
- Angular services introduction
- Angular Dependency Injection (DI)

Student Exercise: 
Create a small angular application that displays an hello world message on the screen
You only have to create a module a controller and a template

#### Angular Templating Engine

The files in this topic are located at the folder angular-template-engine.
In this lesson we will talk about angular templating engine and the main features
in it that are commonly used. We will also talk about these Directives

**Interpolation**
- Calculation in template
- bind controller variable to template
- bind controller function to template

**Architecture Directives**
- ngApp
- ngController
- ngInclude

**Events Directives**
- ngClick
- ngSubmit
- Mouse Events
- ngChange

**Forms**
- Two way binding
- Validation: required, mail validation
- FormControl/NgModelControl
- ***Student Exercise:*** 

Start a new Angular application
In the Index.html add a reference to a controller and a template.
In the template a Form to add a Todo task add the following fields in the form:
title - required and can only contain letters up to 50 length
description - required with minlength of 5
group - required up to length of 20
when - date field
The submit button of the form should be disabled until all values are valid
Attach an event in the form submittion and in that event print all the values from the text fields

**DOM Manipulation**
- ngIf
- ngRepeat

**Changine appearance**
- ngClass
- ngStyle

**Filters**
- dateFilter
- filterFilter
- jsonFilter
- Writing a Custom Filter

**Student Summary Exercise**
- goto the server task list url: https://nztodo.herokuapp.com/api/task/?format=json
- copy all the tasks and print them in a controller property in the controller you created in the task before.
- using ngRepeat print all the tasks on the screen
- in each task add a button to delete that task
- add also a search box at the top of the list and create a search box that will filter the below list 

## Day 2

#### Angular Services

The files for our services demo are located at the folder: **angular-services**
Services represet the Model part of our angular application and in this section we will cover the 
following topics: 

- constant
- factory
- service

**Student Services Exercise**

- Create a constant holding the server url: https://nztodo.herokuapp.com/
- Create a factory which declares a **Task** class which has the following private properties: 
    * id : number
    * title: string
    * description : string
    * group : string
    * when : Date
- Create getters and setters for those private properties
- Create the following methods in your class:
    * toDict - returns a dictionary representation of the object
    * toString - returns a json string of the object
    * fromDict - init the members from a dictionary object
- Create another factory or service called **TaskService** which has the following methods for now signature only no implementation
    * addTask - gets as an argument an instance of a Task
    * updateTask - gets an instance of a task
    * deleteTask - gets an instance of a task
    * getTasks - returns all the tasks
    * getTask - gets a number and returns a single task
    
#### Working with a REST Server

The files in this topic are located at the folder: **angular-rest-server**
In this lesson we are going to cover the following topics

- What is a REST Server
- $http
- Restangular

#### Node Express Heroku demo

Small demonstration of creating a node express server and hosting the server on heroku
The files in this demo are located at the folder **node-express**

## Day 3

#### Angular Routing

In this lesson we will learn how to work with the ui.router package to create routing in our angular application
The files in this demo are located at the folder **angular-routing**

#### Custom Directives

In this lesson we will learn how to create components in angular using Directives, the files in this demo are located at 
the folder **angular-directives**

## Day 4

#### Angular Summary Exercise

In this lesson we will create a final exercise summarizing Angular material that we covered. 
The task solution is located at the folder **summary-task**
In the following task you will have to create a todo application with two screens: 
- The first screen will have two components, one to create or modify a task and one to display the list of tasks
- The second screen will have a single task details.

You task involves the following:
- Create a new angular application
- install the ui.router module
- install the Restangular module
- define 2 routes in your app one for the homepage and one for a task page
- in the homepage create a controller and a template
- the homepage route should have a resolve that will load when the tasks return from the server
- create 2 directives one for creating or modifying a task which will be a form with the inputs of a task
and also an optional input of id, if the id is present the form will send a request to update the task if not the form will create
a new task. Make sure you have a validation in the form
- the second directive will display the list of tasks with their title and description and a delete button.
- when clicking a task it will move to the route of the task page.
- the task page will have a resolve and the page is loaded when the task is grabbed from the server
- the task page will display all the details of the task


- Create another route for you url that