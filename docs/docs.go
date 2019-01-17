// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-17 21:30:14.479809 +0100 CET m=+0.044539006

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "user-microservice is a set of API to manage users",
        "title": "Swagger user-microservice",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Thomas Poignant",
            "url": "https://github.com/thomaspoignant/user-microservice/"
        },
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
        "/health": {
            "get": {
                "description": "health check endpoint to know if the service is up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.healthCheck"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "test swagger",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bottle ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.user"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.healthCheck": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "API return code",
                    "type": "string",
                    "example": "SUCCESS"
                },
                "health": {
                    "description": "Health status of the service",
                    "type": "string",
                    "example": "RUNNING"
                }
            }
        },
        "api.user": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "API return code",
                    "type": "string",
                    "example": "SUCCESS"
                },
                "user": {
                    "description": "informations of a user",
                    "type": "object",
                    "$ref": "#/definitions/entity.User"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "creation date of the entry (format example 2019-01-17T21:03:08.373394+01:00)",
                    "type": "string",
                    "example": "2019-01-17T21:03:08.373394+01:00"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "description": "id of the User (format UUID)",
                    "type": "string",
                    "example": "8da8adc3-0ae9-47b2-884c-ee41e691ff57"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "updated_at": {
                    "description": "last update date of the entry (format example 2019-01-17T21:03:08.373394+01:00)",
                    "type": "string",
                    "example": "2019-01-17T21:03:08.373394+01:00"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
