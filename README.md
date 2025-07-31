# Pulse API

Welcome to **Pulse**, a dynamic social network where users can create their own blogs, write engaging posts on a variety of topics, and share their thoughts with the world.

This application was developed as part of the technical challenge for the PROD '24 contest. For more details, you can find the task description and the OpenAPI specification in the repository.

## Getting Started

Follow these steps to run the application:

1. **Clone the Repository**:
    ```
    git clone https://github.com/realdanielursul/pulse-api
    ```

2. **Install Required Go Packages:**
    ```bash
    go mod tidy
    ```

3. **Run the Application:**
    ```bash
    docker compose up
    ```

4. **Access the API:**
The API will be available at http://localhost:8080/swagger.

## Technologies Used:
- **Gin**: A web framework for building the API.

- **JWT**: For authorization and authentication.

- **PostgreSQL**: A relational database for data storage.

- **Docker**: For app containerization.

- **Swagger**: API documentation for easy reference.

