# Dynamic Profile Endpoint

This project is a simple RESTful API endpoint that returns profile information along with a dynamic cat fact fetched from an external API.

## Task Requirements

- **Endpoint:** `GET /me`
- **Content-Type:** `application/json`
- **Dynamic Data:** Fetches a new cat fact from the [Cat Facts API](https://catfact.ninja/fact) on every request.

### Response Structure

```json
{
  "status": "success",
  "user": {
    "email": "<your email>",
    "name": "<your full name>",
    "stack": "<your backend stack>"
  },
  "timestamp": "<current UTC time in ISO 8601 format>",
  "fact": "<random cat fact from Cat Facts API>"
}
```

## Setup and Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher) installed on your machine.

### Dependencies

This project uses the following Go modules:

- `github.com/gin-gonic/gin`: A popular Go web framework.
- `github.com/gin-contrib/cors`: Gin middleware for CORS.

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/CyberwizD/Dynamic-Profile-Endpoint.git
    cd Dynamic-Profile-Endpoint
    ```

2.  **Install dependencies:**
    The Go modules system will automatically download the required dependencies when you build or run the project. To install them manually, you can run:
    ```bash
    go mod tidy
    ```

## Running the Application Locally

To run the application locally, use the following command:

```bash
go run main.go
```

The server will start on `http://localhost:8080`. You can access the endpoint at `http://localhost:8080/me`.

## Deployment on DigitalOcean

This application is a self-contained Go executable and can be easily deployed to a DigitalOcean droplet.

1.  **Build the binary:**
    To create a production build, run the following command. This will create a binary named `dynamic-profile-endpoint` in your project directory.
    ```bash
    go build -o dynamic-profile-endpoint
    ```

2.  **Upload the binary to your droplet:**
    You can use `scp` to upload the binary to your DigitalOcean droplet.
    ```bash
    scp dynamic-profile-endpoint user@your_droplet_ip:~/
    ```

3.  **Run the application on your droplet:**
    SSH into your droplet and run the binary. You might want to use a process manager like `systemd` or `supervisor` to run the application as a service.
    ```bash
    ./dynamic-profile-endpoint
    ```

## API Endpoint

- **GET /me**

  Returns the dynamic profile information.

  **Example Response:**

  ```json
  {
    "status": "success",
    "user": {
      "email": "wisdomce19@gmail.com",
      "name": "Wisdom Udoye",
      "stack": "Go/Gin"
    },
    "timestamp": "2025-10-17T00:00:00.000000000Z",
    "fact": "A cat can travel at a top speed of approximately 31 mph (49 km) over a short distance."
  }
  ```

## Features

- **CORS:** The API has Cross-Origin Resource Sharing (CORS) enabled to allow requests from any origin.
- **Logging:** The API logs requests and errors for debugging purposes.

## Environment Variables

This task does not require any environment variables.
