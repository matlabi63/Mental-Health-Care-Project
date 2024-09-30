Mental Health Care

This Application is used to manage stress and anxiety, talk to psychologist or mental health specialist and also road information about mental health.

## Tools Used

- Programming language: [Go](https://go.dev/)
- Database: [MySQL](https://www.mysql.com/)
- Web framework: [Echo](https://echo.labstack.com/)
- ORM library: [GORM](https://gorm.io/)
- Request validation: [Validator](https://github.com/go-playground/validator)
- Application configuration: [Viper](https://github.com/spf13/viper)

## Notes

There are two branches in this repository:

- master: REST API application with clean architecture.
- mvc: REST API application with MVC architecture.

## How to Use

1. Clone this repository.

2. Copy the .env file for database configuration. Then, configure the database connection inside that file.

3. Create a new database.

sql
CREATE DATABASE mental_health_care;

4. Run the application.

sh
go run main.go

## Running with Docker

The application can be run as a Docker container with this command. Make sure to adjust the volume setting inside docker-compose.yml file.

Run the application.

sh
docker compose up -d
