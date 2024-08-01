# Task Management API Documentation
  ## Overview
  
  This API provides functionalities to manage tasks, including creating, updating, deleting, and retrieving tasks. It uses an in-memory database for data storage.
  
  ## Requirements
  
  ### Endpoints
  
  - **GET /tasks**
    - Get a list of all tasks.
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        [
          {
            "id": "1",
            "title": "Task 1",
            "description": "Description for Task 1",
            "due_date": "2024-07-31T00:00:00Z",
            "status": "pending"
          },
          ...
        ]
        ```
  
  - **GET /tasks/:id**
    - Get the details of a specific task.
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        {
          "id": "1",
          "title": "Task 1",
          "description": "Description for Task 1",
          "due_date": "2024-07-31T00:00:00Z",
          "status": "pending"
        }
        ```
  
  - **POST /tasks**
    - Create a new task.
    - **Request**:
      - Body:
        ```json
        {
          "title": "New Task",
          "description": "Description for the new task",
          "due_date": "2024-08-01T00:00:00Z",
          "status": "pending"
        }
        ```
    - **Response**: 
      - Status: `201 Created`
      - Body: 
        ```json
        {
          "id": "2",
          "title": "New Task",
          "description": "Description for the new task",
          "due_date": "2024-08-01T00:00:00Z",
          "status": "pending"
        }
        ```
  
  - **PUT /tasks/:id**
    - Update a specific task.
    - **Request**:
      - Body:
        ```json
        {
          "title": "Updated Task",
          "description": "Updated description",
          "due_date": "2024-08-02T00:00:00Z",
          "status": "completed"
        }
        ```
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        {
          "id": "1",
          "title": "Updated Task",
          "description": "Updated description",
          "due_date": "2024-08-02T00:00:00Z",
          "status": "completed"
        }
        ```
  
  - **DELETE /tasks/:id**
    - Delete a specific task.
    - **Response**: 
      - Status: `204 No Content`
  
  ### Error Handling
  
  - **Invalid Request**: 
    - Status: `400 Bad Request`
    - Response: 
      ```json
      {
        "error": "Invalid input"
      }
      ```
  
  - **Resource Not Found**: 
    - Status: `404 Not Found`
    - Response: 
      ```json
      {
        "error": "Task not found"
      }
      ```
  
  ## Instructions
  
  ### Development Environment
  
  1. **Install Go**: Ensure Go is installed on your system.
  2. **Setup Project**:
     ```sh
     mkdir task_manager
     cd task_manager
     go mod init task_manager
  3. **Folder Structure**:
     
  task_manager/
  ├── main.go
  ├── controllers/
  │   └── task_controller.go
  ├── models/
  │   └── task.go
  ├── data/
  │   └── task_service.go
  ├── router/
  │   └── router.go
  ├── docs/
  │   └── api_documentation.md
  └── go.mod
  - **main.go:** Entry point of the application.
  - **controllers/task_controller.go:** Handles incoming HTTP requests.
  - models/task.go: Defines the Task struct.
  - ***data/task_service.go:*** Contains business logic and data manipulation functions.
  - ***router/router.go:*** Sets up the routes and initializes the Gin router.
  - ***docs/api_documentation.md:*** API documentation
  4. **Run the Application**:
  
      To run the application, open a terminal or command prompt and navigate to the `task_manager` directory. Then, execute the following command:
  
      ``` go run main go ```
  
     [API Documentation](https://documenter.getpostman.com/view/36018169/2sA3kdAck4)





























































































































































