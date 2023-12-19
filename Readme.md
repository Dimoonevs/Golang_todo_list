# README for Monolithic Golang Application

## Project Overview

This is a monolithic application providing an API for managing TODO lists. The application is built with Golang and interacts with a PostgreSQL database for data storage.

## Features

### TODO Item
- **getAllItems:** Retrieve a list of all TODO items.
- **createItem:** Create a new TODO item.
- **getItemById:** Get information about a TODO item by its identifier.
- **updateItem:** Update information about a TODO item.
- **deleteItem:** Delete a TODO item.

### TODO List
- **getAllIList:** Retrieve a list of all TODO lists.
- **createList:** Create a new TODO list.
- **getListById:** Get information about a TODO list by its identifier.
- **updateList:** Update information about a TODO list.
- **deleteList:** Delete a TODO list.



### Authentication and Authorization
- Methods implemented for user registration and authentication using JWT (JSON Web Tokens) from the `github.com/dgrijalva/jwt-go` library.

## Project Structure

- **handlers:** Contains request handlers, including handlers for registration, authentication, and handlers for working with TODO items and lists.
- **services:** Service layer for handling business logic.
- **repositories:** Repository layer for interacting with the PostgreSQL database.

## Technologies

- **Golang:** Programming language for application development.
- **PostgreSQL:** Relational database for data storage.
- **JWT:** Library for working with JSON Web Tokens.
- **sqlx:** Enhanced database/sql library, `github.com/jmoiron/sqlx`.
- **godotenv:** A library for loading environment variables from a file, `github.com/joho/godotenv`.
- **pq:** PostgreSQL driver for Go's database/sql package, `_ "github.com/lib/pq"`.
- **logrus:** Structured logger for Go (logging), `github.com/sirupsen/logrus`.
- **viper:** Configuration library for Go applications, `github.com/spf13/viper`.

## Installation and Configuration

1. **Install Golang:**
   Follow the [official documentation](https://golang.org/doc/install) to install Golang.

2. **Install PostgreSQL:**
   Install PostgreSQL and create a database. Use the configuration files in the `config` folder to configure the connection.

3. **Install Dependencies:**
   In the project's root directory, run:
   ```bash
   go mod download
   ```

## Running the Application

1. **Start PostgreSQL:**
   Ensure that PostgreSQL is running, and the database is created.

2. **Run the Application:**
   In the project's root directory, run:
   ```bash
   go run main.go
   ```

## Notes

Be sure to change the database connection parameters in the configuration files in the `config` folder to match your configuration.