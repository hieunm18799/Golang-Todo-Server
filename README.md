# Todo Server - Golang RESTful API Project

Creating a RESTful API with authentication to perform the basic todo's Create/Get/Update/Delete operations

# Getting Started

## Dependencies

- The following feature set is a minimal selection of typical Web API requirements:

    - Configuration using [viper](https://github.com/spf13/viper)
    - PostgreSQL support including migrations using [postgres](https://github.com/go-gorm/postgres)
    - Routing with [gin](https://github.com/gin-gonic/gin) and middleware
    - JWT Authentication using [jwx/v5](https://github.com/golang-jwt/jwt)
    - [ORM](https://gorm.io/) for Golang
    - Google's [uuid](https://github.com/google/uuid) for unique id generation

- Command line to get all features:
    ```bash
    go get - u github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres github.com/spf13/viper github.com/google/uuid
    ```

## Usage

- Clone this repository
- Create "app.env" file in the main directory for environment variables

    - `POSTGRES_HOST`=127.0.0.1
    - `POSTGRES_USER`=postgres
    - `POSTGRES_PASSWORD`=<>
    - `POSTGRES_DB_NAME`=golang
    - `POSTGRES_PORT`=6500
    - `PORT`=8000
    - `CLIENT_ORIGIN`=http://localhost:3000

- Create docker's compose file in the main directory to manage an instance of PostgreSQL(example)
    ```bash
    services:
    postgres:
        image: postgres
        container_name: todos
        ports:
        - 6500:5432
        env_file:
        - ./app.env
        volumes:
        - postgres: <docker volumes>
    volumes:
    postgres:
    ```

- Run the docker and application with commands:
    ```bash
    docker-compose up -d
    go run  .\cmd\server\main.go
    ```

## API Routes

| Path                         | Method | Header               | Required JSON    | Description                                 |
|------------------------------|--------|----------------------|------------------|---------------------------------------------|
| /api/user/register           | POST   |                      | user             | register user and return user               |
| /api/token                   | POST   |                      | email + password | get jwt by email and password               |
| /api/secured/ping            | GET    | Authorization: "JWT" |                  | just for test authentication                |
| /api/secured/todos           | GET    | Authorization: "JWT" |                  | get todos depend on the user's              |
| /api/secured/todo/:id        | GET    | Authorization: "JWT" |                  | get todo depend on the user's               |
| /api/secured/todos           | POST   | Authorization: "JWT" | todo             | create todo and return it                   |
| /api/secured/todo/:id        | PATCH  | Authorization: "JWT" | todo             | Fix todo (not ID) if it belongs to the user |
| /api/secured/todo_change/:id | DELETE | Authorization: "JWT" |                  | Delete todo if it belongs to the user       |