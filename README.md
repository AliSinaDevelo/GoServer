# Simple Go Web Server

This is a basic web server written in Go that handles simple form submissions and provides a "hello" endpoint. The server serves static files from the `static` directory.

## Features

- **Form Handling:** Accepts POST requests to the `/form` endpoint and displays the submitted data.
- **Hello Endpoint:** Responds with a simple "Howdy!" message for GET requests to the `/hello` endpoint.
- **Static File Serving:** Serves static files from the `static` directory.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/AliSinaDevelo/your-repo.git
   cd your-repo

2. Run the server:
   
   ```bash
   go run main.go
  

3. Access the following endpoints in your web browser or through a tool like curl or httpie:

  http://localhost:8080/form: Form handling endpoint
  http://localhost:8080/hello: Hello endpoint

## Dependencies
  Go

## License
  This project is licensed under the MIT License.

