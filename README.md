# Pulse API

Welcome to **Pulse**, a dynamic social network where users can create their own blogs, write engaging posts on a variety of topics, and share their thoughts with the world.

This application was developed as part of the technical challenge for the PROD '24 contest. For more details, you can find the task description and the OpenAPI specification in the `docs` directory.

## Getting Started

Follow these steps to run the application locally:

1. **Clone the Repository**:
    ```
    git clone https://github.com/realdanielursul/pulse-api
    ```

2. **Install Required Go Packages:**
    ```bash
    go mod tidy
    ```

3. **Configure the application:**
    - Configure `config.yaml` file in `configs` directory.
    - Create `.env` file in root directory and set environment variables as following:

        ```
        DB_PASSWORD=...
        SECRET_KEY=...
        SALT=...
        ```

4. **Run the Application:**
    ```bash
    make up && make run
    ```

5. **Access the API:**
The API will be available at http://localhost:8080/swagger.
