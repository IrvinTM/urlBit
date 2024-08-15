
# URLBIT 
UrlBit is a URL shortener API using Go, PostgreSQL, and JSON Web Tokens for authentication.

## Prerequisites 
- Docker
- Docker Compose

## Deploy
Modify the **docker-compose** file, set the database credentials as you wish, and add the JWT secret.

Start the containers:
```bash
docker compose up -d
```

## API Endpoints
Below are the available API endpoints for this project:

### Authentication Endpoints

- Register a new account<br/>
  **Path:** `/api/register`  
  **Method:** `POST`  
  **Description:** Creates a new user account.  
  **Body Parameters:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- Login to an existing account<br/>
  **Path:** `/api/login`  
  **Method:** `POST`  
  **Description:** Authenticates a user and returns a token.  
  **Body Parameters:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

### URL Management Endpoints

- Create a new shortened URL<br/>
  **Path:** `/api/newurl`  
  **Method:** `POST`  
  **Description:** Creates a new shortened URL.  
  **Body Parameters:**
  ```json
  {
    "original_url": "string"
  }
  ```

- Get all URLs for the authenticated user<br/>
  **Path:** `/api/myurls`  
  **Method:** `GET`  
  **Description:** Retrieves all URLs created by the authenticated user.  
  **Response:**
  ```json
  [
    {
      "short_url": "string",
      "original_url": "string"
    }
  ]
  ```

### URL Redirection

- Redirect to the original URL and update click count<br/>
  **Path:** `/{shorturl}`  
  **Method:** `GET`  
  **Description:** Redirects the user to the original URL based on the shortened URL provided.