# Mental Health Care

This application serves as a comprehensive platform for managing stress and anxiety. It connects users with qualified doctor and mental health specialists for personalized support while also offering a wealth of resources and information to promote mental well-being. Whether you're seeking professional guidance or looking to educate yourself about mental health, this application is your go-to resource for achieving a healthier, more balanced mind.

## Tools Used

- Programming language: [Go](https://go.dev/)
- Database: [MySQL](https://www.mysql.com/)
- Web framework: [Echo](https://echo.labstack.com/)
- ORM library: [GORM](https://gorm.io/)
- Request validation: [Validator](https://github.com/go-playground/validator)
- Application configuration: [Viper](https://github.com/spf13/viper)

- ## Mental_Health_care ERD: [Mental Health Care Project Diagram](https://www.drawdb.app/editor?shareId=7d20770ad62c4c45d34e18eb7b348244)

- ## Mental_Health_care UseCase: [Mental Health Care Project UseCase](https://drive.google.com/file/d/11JkDkyQBQFAfTwE7VD3i58ER_2mfqW1O/view?usp=sharing)

- ## API Documentation: [Mental Health Care Application](https://mental-health-care.postman.co/workspace/Mental-Health-Care-Workspace~346a1053-7058-45d0-9788-c214bb57e3c3/collection/36680903-82950e84-d953-462d-ad3b-17eebee21c63?action=share&creator=36680903)

## Key Features

There are two branches in this repository:

- master: REST API application with clean architecture.
- mvc: REST API application with MVC architecture.

- There are some key features in this application;

### Users:

- Users can register for an account, securely log in, submit complaints, and access valuable recommendations for managing their mental health.

### Admins:

- Admins have the authority to manage both end users and doctors, as well as oversee content within the application.

### Doctors:

- Doctors can provide recommendations to end users regarding mental health.

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

The application can be run as a Docker container with this command. Make sure to adjust the volume setting inside docker file.

Run the application.

sh
docker compose up -d
