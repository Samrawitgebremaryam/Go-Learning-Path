# Task Management REST API Documentation

## Introduction

The Task Management REST API facilitates the management of tasks through CRUD operations. It uses JWT for secure user authentication and authorization.

## Base URL

```
http://localhost:8080
```

## Endpoints

### Authentication

#### Register User

**POST /register**

**Description:** Registers a new user.

**Request:**

```json
{
  "email": "user@example.com",
  "password": "password123",
  "usertype": "Admin" // or "User"
}
```

**Response:**

- **200 - OK**

```json
{ "message": "User registered successfully" }
```

#### Login User

**POST /login**

**Description:** Authenticates a user and returns a JWT token.

**Request:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**

- **200 - OK**

```json
{
  "message": "User logged in successfully",
  "token": "jwt_token_here"
}
```

### Tasks

All task endpoints require the `Authorization` header with the value `Bearer <token>`.

#### Create a Task (Admin Only)

**POST /tasks**

**Description:** Creates a new task.

**Request Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request:**

```json
{
  "title": "Prepare",
  "description": "Prepare the slides.",
  "due_date": "2025-07-10",
  "status": "In Progress"
}
```

**Response:**

- **201 - Created**

```json
{
  "message": "Task created successfully",
  "task": {
    /* task details */
  }
}
```

#### Get All Tasks

**GET /tasks**

**Description:** Retrieves all tasks.

**Request Headers:**

```
Authorization: Bearer <token>
```

**Response:**

- **200 - OK**

```json
[
  /* list of tasks */
]
```

#### Get a Specific Task

**GET /tasks/:id**

**Description:** Retrieves a specific task by ID.

**Request Headers:**

```
Authorization: Bearer <token>
```

**Response:**

- **200 - OK**

```json
{
  /* task details */
}
```

- **404 - Not Found**

```json
{ "error": "Task not found" }
```

#### Update a Task

**PUT /tasks/:id**

**Description:** Updates a specific task by ID.

**Request Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request:**

```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-07-15",
  "status": "Completed"
}
```

**Response:**

- **200 - OK**

```json
{
  "message": "Task updated successfully",
  "task": {
    /* updated task details */
  }
}
```

- **404 - Not Found**

```json
{ "error": "Task not found" }
```

#### Delete a Task (Admin Only)

**DELETE /tasks/:id**

**Description:** Deletes a specific task by ID.

**Request Headers:**

```
Authorization: Bearer <token>
```

**Response:**

- **200 - OK**

```json
{ "message": "Task deleted successfully" }
```

- **404 - Not Found**

```json
{ "error": "Task not found" }
```

## Testing with Postman

### Setup Environment

Ensure the API is running locally on `http://localhost:8080`.

### Create Requests

- **Register User:** POST to `http://localhost:8080/register` with JSON body.
- **Login User:** POST to `http://localhost:8080/login` with JSON body to get the token.
- **Authenticated Requests:** Add `Authorization: Bearer <token>` header.

### Example Requests

1. **Register User:**

   - Method: POST
   - URL: `http://localhost:8080/register`
   - Body:
     ```json
     {
       "email": "user@example.com",
       "password": "password123",
       "usertype": "Admin"
     }
     ```

2. **Login User:**

   - Method: POST
   - URL: `http://localhost:8080/login`
   - Body:
     ```json
     {
       "email": "user@example.com",
       "password": "password123"
     }
     ```

3. **Create Task (Admin Only):**

   - Method: POST
   - URL: `http://localhost:8080/tasks`
   - Headers: `Authorization: Bearer <token>`
   - Body:
     ```json
     {
       "title": "Prepare",
       "description": "Prepare the slides.",
       "due_date": "2025-07-10",
       "status": "In Progress"
     }
     ```

4. **Get All Tasks:**

   - Method: GET
   - URL: `http://localhost:8080/tasks`
   - Headers: `Authorization: Bearer <token>`

5. **Get Specific Task:**

   - Method: GET
   - URL: `http://localhost:8080/tasks/1`
   - Headers: `Authorization: Bearer <token>`

6. **Update Task:**

   - Method: PUT
   - URL: `http://localhost:8080/tasks/1`
   - Headers: `Authorization: Bearer <token>`
   - Body:
     ```json
     {
       "title": "Updated Title",
       "description": "Updated Description",
       "due_date": "2025-07-15",
       "status": "Completed"
     }
     ```

7. **Delete Task (Admin Only):**
   - Method: DELETE
   - URL: `http://localhost:8080/tasks/1`
   - Headers: `Authorization: Bearer <token>`

### Conclusion

This concise documentation covers user registration, login, and task management with JWT-based authentication and authorization.

for more info : https://documenter.getpostman.com/view/37193879/2sA3kdBHnP
