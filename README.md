# Gin-API Project

Welcome to the API Project! This project is built using PostgreSQL and the Gin framework for Go and provides a simple RESTful API with basic authentication and document management features.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker](#docker)
- [Contributing](#contributing)
- [License](#license)

## Features

- Sign-In: Authenticate users with their credentials.
- Sign-Up: Register new users to the system.
- Create Document: Allow authenticated users to create and manage documents.


## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/KarmaBeLike/Gin-Api
    cd the repository
    ```

2. Install dependencies:
    ```bash
    go mod download
    ```



## Usage

1. Run the application:
    ```bash
    go run ./cmd
    ```

2. Access the application at `http://localhost:8080`.

## API Endpoints

### Authentication and Session

- **Sign Up**
    - `POST /signup`
    - Description: Create a new user account.
    - Request Body: 
      ```json
      {
          "username": "string",
          "password": "string"
      }
      ```
    - Response: 
      ```json
      {
          "message": "User succcessfully registered"
      }
      ```

- **Sign In**
    - `POST /signin`
    - Description: Authenticate a user and start a session.
    - Request Body: 
      ```json
      {
          "username": "string",
          "password": "string"
      }
      ```
    - Response: 
      ```json
      {
          "message": "User successful logged in"

      }
      ```

### Documents

- **Get Document**
    - `GET /getdoc`
    - Description: Retrieve a document.
    - Response: 
      ```json
      {
          "document": {
              "id": "int",
              "title": "string",
              "content": "string"
          }
      }
      ```

- **Create Document**
    - `POST /documents`
    - Description: Create a new document.
    - Request Body: 
      ```json
      {
          "title": "string",
          "content": "string"
      }
      ```
    - Response: 
      ```json
      {
          "message": "Document created successfully"
      }
      ```


