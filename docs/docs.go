// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/task": {
            "get": {
                "description": "Get a sample response from the server",
                "produces": [
                    "application/json"
                ],
                "summary": "GetTask",
                "operationId": "get-task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Resource ID",
                        "name": "task_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resource found",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            },
            "put": {
                "description": "Get a sample response from the server",
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateTaskStatus",
                "operationId": "put-task-status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Resource ID",
                        "name": "task_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resource updated successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Get a sample response from the server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateAndUpdateTask",
                "operationId": "post-task",
                "parameters": [
                    {
                        "description": "Request Body if send ID that will become UPDATE API",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resource created successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Get a sample response from the server",
                "produces": [
                    "application/json"
                ],
                "summary": "DeleteTask",
                "operationId": "put-task-status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Resource ID",
                        "name": "task_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resource deleted successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    }
                }
            }
        },
        "/task-list": {
            "get": {
                "description": "Get a sample response from the server",
                "produces": [
                    "application/json"
                ],
                "summary": "GetTasksList",
                "operationId": "get-task-list",
                "responses": {
                    "200": {
                        "description": "Resource found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Task": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.TaskRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/task-management/api/v1",
	Schemes:          []string{},
	Title:            "Task Management API",
	Description:      "This is a sample Echo API with Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}