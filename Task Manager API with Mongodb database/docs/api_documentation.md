# Task Management REST API Documentation

## Introduction

The Task Management REST API is designed to facilitate the management of tasks through a series of CRUD (Create, Read, Update, Delete) operations. Developed using the Go programming language and the Gin web framework, this API allows users to seamlessly interact with tasks, enabling them to create new tasks, retrieve existing ones, update task details, and delete tasks as needed.

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Create a Task

**POST /tasks**

**Description:** Creates a new task with the provided details such as title, description, due date, and status.

**URL:** `/tasks`

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

```json
{
  "id": 6,
  "title": "Prepare",
  "description": "Prepare the slides.",
  "due_date": "July 10, 2025",
  "status": "In Progress"
}
```

**Response:**

- **202 - Accepted**

```json
{
  "message": "Task created successfully",
  "task": {
    "id": 6,
    "title": "Prepare",
    "description": "Prepare the slides.",
    "due_date": "July 10, 2025",
    "status": "In Progress"
  }
}
```

---

### 2. Get All Tasks

**GET /tasks**

**Description:** Retrieves a list of all tasks stored in the system.

**URL:** `/tasks`

**Response:**

- **200 - OK**

```json
[
  {
    "id": 1,
    "title": "Complete Project Proposal",
    "description": "Draft and finalize the project proposal for client approval.",
    "due_date": "July 12, 2024",
    "status": "In Progress"
  },
  {
    "id": 2,
    "title": "Implement Authentication",
    "description": "Develop and test user authentication features.",
    "due_date": "July 13, 2024",
    "status": "Not Started"
  }
]
```

---

### 3. Get a Specific Task

**GET /tasks/:id**

**Description:** Retrieves the details of a specific task identified by its unique ID.

**URL:** `/tasks/:id`

**Path Variables:**

- `id` (number): The unique identifier of the task.

**Response:**

- **200 - OK**

```json
{
  "id": 1,
  "title": "Complete Project Proposal",
  "description": "Draft and finalize the project proposal for client approval.",
  "due_date": "July 12, 2024",
  "status": "In Progress"
}
```

- **404 - Not Found**

```json
{
  "error": "Task not found"
}
```

---

### 4. Update a Task

**PUT /tasks/:id**

**Description:** Updates the details of a specific task identified by its unique ID. The request body should include the fields to be updated.

**URL:** `/tasks/:id`

**Path Variables:**

- `id` (number): The unique identifier of the task to be updated.

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "July 15, 2025",
  "status": "Completed"
}
```

**Response:**

- **200 - OK**

```json
{
  "message": "Task updated successfully",
  "task": {
    "id": 1,
    "title": "Updated Title",
    "description": "Updated Description",
    "due_date": "July 15, 2025",
    "status": "Completed"
  }
}
```

- **404 - Not Found**

```json
{
  "error": "Task not found"
}
```

---

### 5. Delete a Task

**DELETE /tasks/:id**

**Description:** Deletes a specific task identified by its unique ID.

**URL:** `/tasks/:id`

**Path Variables:**

- `id` (number): The unique identifier of the task to be deleted.

**Response:**

- **200 - OK**

```json
{
  "message": "Task deleted successfully"
}
```

- **404 - Not Found**

```json
{
  "error": "Task not found"
}
```

---

## Testing with Postman

### Setup Environment

Ensure the API is running locally on `http://localhost:8080`.

### Create Requests

For each endpoint, create a new request in Postman:

- **POST /tasks:** Use the POST method, set the URL to `http://localhost:8080/tasks`, and include the JSON body as specified.
- **GET /tasks:** Use the GET method, set the URL to `http://localhost:8080/tasks`.
- **GET /tasks/:id:** Use the GET method, set the URL to `http://localhost:8080/tasks/1` (replace `1` with the actual task ID).
- **PUT /tasks/:id:** Use the PUT method, set the URL to `http://localhost:8080/tasks/1` (replace `1` with the actual task ID), and include the JSON body as specified.
- **DELETE /tasks/:id:** Use the DELETE method, set the URL to `http://localhost:8080/tasks/1` (replace `1` with the actual task ID).

### Send Requests

Send each request and verify the response matches the expected format and status code.

### Document Results

For each request, document the request details, including headers, body, and the response received. Save these in Postman collections for future reference.

---
