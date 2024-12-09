# Golang CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API built with Go (Golang) using the Gin framework. The API manages tasks and supports authentication via JWT. It stores task data in a PostgreSQL database and provides Swagger documentation for easy testing.

## Setup and Run (Guide)

### Prerequisites
- Go (Golang) installed on your machine (version 1.18 or higher)
- PostgreSQL installed and running
- [golang-migrate](https://github.com/golang-migrate/migrate) tool installed for running migrations

### Installation Steps

1. **Clone the Repository**:
   First, clone the repository to your local machine:
   ```bash
   git clone https://github.com/serafinasalim/golang-crud-api.git
   cd golang-crud-api
2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
3. **Set Up Database**:
   - **Create Database**: If you don’t have the database created yet, run the following commands in your PostgreSQL terminal 
      ```sql 
      CREATE DATABASE dbname;
      ```
   - **Migrate Database**:\
      Make sure migrate is installed by checking with Command Prompt/Terminal
      ```bash
      migrate -version
      ```
      Then navigate to the cloned directory
      ```bash
      cd PathToClonedDirectory
      ```
      Run Migration with the command below (replace user, password, and dbname with yours)
      ```bash
      migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
      ```
   - **Change Database Const**:\
      Navigate to databse/db.go in the repository and change the value to your database
      ```go
     const (
	      DBUser     = "user"
	      DBPassword = "password"
	      DBName     = "dbname"
      )
      ```
4. **Run the Application**:
   ```bash
   go run .
   ```
5. **Get JWT Token**:\
   To test protected endpoints you need to get JWT Token by hitting Register and Login public endpoints. Then authenticate in the provided field in Swagger to test Tasks Endpoint.

This will start the application locally and you can test it in provided 
[Swagger Documentation](http://localhost:8080/docs/swagger/index.html)