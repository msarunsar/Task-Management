# Task Management
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
## Requirements
- Go (version 1.16 or later)
- Echo framework (installed via `go get -u github.com/labstack/echo/v4`)
- Echo swagger (installed via `go install github.com/swaggo/swag/cmd/swag@latest`)
## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
2. Run serve:
   ```bash
   go mod tidy
   go run main.go
## Usage
- server running on: http://localhost:8080/
- swagger running on: http://localhost:8080/swagger/index.html
## API Endpoints
| NAME   |      METHOD      | ENDPOINT | REQUEST | REMARK |
|----------|:-------------:|------|:------:|------|
| create or update task | POST | task-management/api/v1/task | body | if body request have **ID** that will become **UPDATE** |
| get task by id |    GET   |   task-management/api/v1/task | param |  |
| get task list |    GET   |   task-management/api/v1/task-list | N/A |  |
| update task status |    PUT   |   task-management/api/v1/task | param |  |
| delete task |    DELETE   |   task-management/api/v1/task | param |  |

more detail in document file: ***"Task Management API.pdf"***