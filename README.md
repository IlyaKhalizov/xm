### Build & Deploy

Project uses postgres sql data base dependency with migrations

Before a deploy, set proper env file according to envexample

Execute the command docker-compose up -d
    - builds go image container
    - builds postgres db container
    - creates migrations

### Make

Before build, run make file (lint and test)

Make migration_up and migration_down create database migrations or eliminate them

### API

POST /v1/company - creates new company

    {
        "Name": "Test2",
        "Description": "Test123",
        "EmployeesAmount": 10,
        "Registered": true,
        "Type": "corporations"
    }

GET /v1/company/{id} - gets a company by id

PATCH /v1/company/{id} - updates a company by id

    {
        "Name": "Test2",
        "Description": "Test123",
        "EmployeesAmount": 10,
        "Registered": true,
        "Type": "corporations"
    }

DELETE /v1/company/{id} - deletes a company by id
