User Age API

Go + Fiber + PostgreSQL + SQLC


Project Overview:


This project is a RESTful API built using Go that manages users with their name and date of birth (DOB).

The API does not store age in the database.
Instead, it calculates age dynamically when user data is requested.

This project follows clean backend architecture and is suitable for real-world applications.


What This Project Demonstrates:


Backend development using Go


REST APIs using GoFiber


Database design using PostgreSQL


Type-safe SQL queries using SQLC


Clean architecture (handler, service, repository)


Input validation


Logging


API testing


Tech Stack:


Go


GoFiber


PostgreSQL


SQLC


Uber Zap (Logging)


Postman (API testing)


Project Structure:


user-age-api/
│
├── cmd/server/main.go        # Application entry point
├── db/
│   ├── migrations/           # DB schema
│   └── sqlc/queries/         # SQLC queries
├── internal/
│   ├── handler/              # HTTP handlers
│   ├── service/              # Business logic (age calculation)
│   ├── repository/           # Database connection
│   ├── routes/               # API routes
│   ├── models/               # Request & response models
│   ├── middleware/           # Logging middleware
│   └── logger/               # Zap logger
├── go.mod
└── README.md



SETUP & RUN:


1. Install Go


Step 1: Download Go


Visit: https://go.dev/dl/


Download Go for Windows


Install with default settings


Step 2: Verify Installation


Open Command Prompt and run:
go version


Expected output:
go version go1.xx windows/amd64


2.Install PostgreSQL


Step 1: Download PostgreSQL


Visit: https://www.postgresql.org/download/windows/


Download installer


During installation:


Remember the password


Keep port as 5432


Step 2: Open SQL Shell (psql)


After installation:


Start Menu → SQL Shell (psql)


Press Enter for defaults


Enter password when asked


3.Create Database & Table


Step 1: Create Database


Inside psql, run:
CREATE DATABASE user_age_db;


Step 2: Connect to Database

\c user_age_db;


Step 3: Create users Table


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);


4.Set Database Environment Variable


The application reads DB credentials from an environment variable.


set DATABASE_URL=postgres://postgres:PASSWORD@localhost:5432/user_age_db?sslmode=disable


If your password contains @, replace it with %40


5.Install Project Dependencies


run :
go mod tidy


6.Run the Application


go run cmd/server/main.go



Expected Output


starting user-age-api server


server running port=3000


Fiber v2.x.x


http://127.0.0.1:3000


API TESTING:


Tool Used


Postman


Download:


https://www.postman.com/downloads/


Base URL:


http://localhost:3000


Create User:


POST /users


Body (JSON):


{
  "name": "xxxxxx",
  "dob": "yyyy-mm-dd"
}


Expected:


Status: 201 Created


Age calculated dynamically


Design Explanation:


Age is not stored in database


Calculated dynamically using Go time package


Service layer contains business logic


Repository layer handles DB access


SQLC provides type-safe queries


Zap logger logs every request


APIs tested using Postman


Testing Summary:


All CRUD APIs tested


Correct HTTP status codes verified


Database operations verified


Age calculation verified