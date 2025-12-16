
# Overview


The goal of this task was to build a REST API in Go that stores user details and calculates age dynamically based on the date of birth.


While implementing the solution, I focused on:


1. keeping the design simple and clean,


2. following standard backend practices,


3. and making the code easy to understand and explain.


This document explains what I built,How I built, why I made certain decisions, and how the system works end to end.


## What I Built

I built a RESTful backend API using Go and GoFiber that supports:


1. Creating a user with name and date of birth


2. Fetching a user by ID


3. Listing all users


4. Updating a user


5. Deleting a user


The database stores only:


1. id


2. name


3. dob (date of birth)


The user’s age is not stored in the database.


Instead, it is calculated dynamically whenever user data is requested.


## Why I Did NOT Store Age in the Database

Age is a value that changes over time.


If age is stored in the database, it becomes incorrect every year.


So instead of storing age:


1. I stored only the date of birth


2. I calculated age at runtime using Go’s time package


This ensures:


1. data consistency


2. no need for background jobs to update age


3. correct age at any point in time


This logic is handled in the service layer.



## How I Built the Project (End-to-End)


### Step 1: Project Setup

I started by:


1. Initializing a Go module


2. Creating the folder structure as given in the assessment


The structure separates:


1. application entry point (main.go)


2. business logic


3. database logic


4. routing


5. middleware


This helped keep the code organized from the beginning.



### Step 2: Database Design


I designed a simple PostgreSQL database with one table:

users(id, name, dob)


I intentionally did not store age because age changes over time.
Instead, only the date of birth is stored.

### Step 3: SQL Queries and SQLC


I wrote SQL queries for:


1. create user


2. get user by ID


3. list users


4. update user


5. delete user


Then I used SQLC to generate type-safe Go code from these SQL queries.


This allowed me to:


1. avoid writing raw SQL in Go files


2. catch query-related issues at compile time


3. keep database logic clean and safe


### Step 4: Repository Layer


I created a repository layer to:


1. establish the database connection


2. configure connection pooling


3. provide SQLC query access to the service layer


This layer only handles database communication and nothing else.


### Step 5: Service Layer 


The service layer is where I placed the core business logic.


Here I:


1. calculated the user’s age dynamically using Go’s time package


2. converted database models into API response models


Age calculation logic checks whether the birthday has already occurred in the current year to ensure accuracy.


### Step 6: Handler Layer 


The handler layer:

1. receives HTTP requests


2. parses JSON input


3. validates request data


4. calls the service layer


5. returns proper HTTP responses


I kept handlers simple and free from business or database logic.


### Step 7: Validation


I used go-playground/validator to validate:


1. required fields


2. correct date format


This prevents invalid data from reaching the database.


### Step 8: Routing


I defined all API routes in a separate routes file.


This keeps:


1. main.go clean


2. routing logic centralized


3. endpoints easy to manage


### Step 9: Logging and Middleware


I added a logging middleware using Uber Zap.


For every request, it logs:


1. HTTP method


2. request path


3. response status


4. request duration


This simulates real production logging practices.


### Step 10: Application Startup (main.go)


In main.go, I:


1. initialized the logger


2. connected to the database using environment variables


3. created service and handler instances


4. registered routes


5. started the Fiber server


This file acts as the entry point that wires everything together.


### Step 11: API Testing


After building the application, I tested all APIs using Postman.


I verified:


1. all CRUD operations

2. correct HTTP status codes


3. dynamic age calculation


4. proper error handling


This confirmed that the application works end to end.


## Project Structure and Design Decisions


I followed a layered architecture to separate responsibilities clearly.


### Handler Layer


1. Handles HTTP requests and responses


2. Parses input JSON


3. Performs input validation


4. Returns proper HTTP status codes


This layer does not contain business logic or database logic.


### Service Layer


1. Contains business logic


2. Calculates user age dynamically


3. Converts database models into API response models


Placing age calculation here keeps the logic reusable and testable.


### Repository Layer


1. Responsible only for database access


2. Manages database connections


3. Uses SQLC-generated queries
   

This separation makes the code easier to maintain and change.


## Why I Used SQLC Instead of an ORM


I used SQLC because:


1. It allows writing plain SQL queries


2. It generates type-safe Go code


3. Errors are caught at compile time


4. There is no hidden ORM behavior


This gives better control over database queries and performance.


## Database Design


The database has a single table:

users(id, name, dob)


The schema is intentionally minimal and follows the task requirements exactly.


## Validation and Error Handling


I used go-playground/validator to:


1. ensure required fields are present


2. validate date format (YYYY-MM-DD)


For errors:


1. invalid input returns 400 Bad Request


2. missing records return 404 Not Found


3. server errors return 500 Internal Server Error


This keeps API behavior predictable and clean.


## Logging


I used Uber Zap for logging.


A logging middleware logs:


1. HTTP method


2. API path


3. response status


4. request duration


This is useful for debugging and monitoring in real applications.


## How the System Works End-to-End


1. Client sends an HTTP request (for example, POST /users)


2. The handler: parses the request and validates input


3. The service: applies business logic (calculates age )


4. The repository: executes SQL queries using SQLC


5. The response is sent back as JSON


## Testing Approach


I tested all APIs using Postman, including:


1. create user


2. fetch user


3. list users


4. update user


5. delete user


I verified:


1. correct status codes


2. correct responses


3. dynamic age calculation










