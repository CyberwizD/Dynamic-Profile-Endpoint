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

- [Go](https://golang.org/dl/) installed on your machine.

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/CyberwizD/Dynamic-Profile-Endpoint.git
    cd Dynamic-Profile-Endpoint
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

## Running the Application

To run the application locally, use the following command:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

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
    "timestamp": "2025-10-16T20:09:30.123456789Z",
    "fact": "A cat can travel at a top speed of approximately 31 mph (49 km) over a short distance."
  }
  ```

## Features

- **CORS:** The API has Cross-Origin Resource Sharing (CORS) enabled to allow requests from any origin.
- **Logging:** The API logs requests and errors for debugging purposes.

## Environment Variables

This task does not require any environment variables.
