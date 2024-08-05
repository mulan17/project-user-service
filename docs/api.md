# RESTful API for user-service

## Resources

- **`/users`**
- **`/users/{id}`**
- **`/admin/block/{id}`**
- **`/admin/limit/{id}`**
- **`/login`**

## Supported Operations
- `POST /users` - Adds a new user.
- `GET /users` - Retrieves information about all users.- - `PATCH /users/{id}` - Updates information about a specific user.
- `GET /users/{id}` - Retrieves information about a specific user by ID. Authentication is required.
- `POST /admin/block/{id}` - Blocks a specific user by ID. Authentication is required.
- `POST /admin/limit/{id}` - Limits a specific user by ID. Authentication is required.
- `POST /login` - Authenticates a user and provides a session or token.

## Methods

### **`/users`**

- **Create**
  - **Method:** `POST`
  - **Description:** Adds a new user.
  - **Example Request:**
    ```http
    POST /users
    Content-Type: application/json

    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```

- **Read All**
  - **Method:** `GET`
  - **Description:** Retrieves information about all users.
  - **Example Request:**
    ```http
    GET /users
    ```

- **Update**
  - **Method:** `PATCH`
  - **Description:** Updates information about a specific user.
  - **Example Request:**
    ```http
    PATCH /users/{id}
    Content-Type: application/json

    {
      "email": "newemail@example.com"
    }
    ```

### **`/users/{id}`**

- **Read by ID**
  - **Method:** `GET`
  - **Description:** Retrieves information about a specific user by ID. Authentication is required.
  - **Example Request:**
    ```http
    GET /users/{id}
    ```

### **`/admin/block/{id}`**

- **Block User**
  - **Method:** `POST`
  - **Description:** Blocks a specific user by ID. Authentication is required.
  - **Example Request:**
    ```http
    POST /admin/block/{id}
    ```

### **`/admin/limit/{id}`**

- **Limit User**
  - **Method:** `POST`
  - **Description:** Limits a specific user by ID. Authentication is required.
  - **Example Request:**
    ```http
    POST /admin/limit/{id}
    ```

### **`/login`**

- **Login**
  - **Method:** `POST`
  - **Description:** Authenticates a user and provides a session or token.
  - **Example Request:**
    ```http
    POST /login
    Content-Type: application/json

    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```
