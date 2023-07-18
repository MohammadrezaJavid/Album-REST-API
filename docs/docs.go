// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/api",
            "email": "api@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "get": {
                "description": "Retrieve all albums from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get album"
                ],
                "summary": "Get all Albums from database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                }
            },
            "put": {
                "description": "Retrieve an Album by id from database, update it and save to database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "put album"
                ],
                "summary": "Update one Album from database",
                "parameters": [
                    {
                        "description": "Album JSON",
                        "name": "album",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                }
            },
            "post": {
                "description": "Post an Album and saved to database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post album"
                ],
                "summary": "Post an Album",
                "parameters": [
                    {
                        "description": "Album JSON",
                        "name": "album",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                }
            }
        },
        "/albums/{id}": {
            "get": {
                "description": "Retrieve an Album by id from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get album"
                ],
                "summary": "Get one Album from database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an Album by id from database.",
                "tags": [
                    "delete album"
                ],
                "summary": "Delete one Album from database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Album": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
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
	Host:             "localhost:8070",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Go + Gin Album rest API Service",
	Description:      "This is a sample server for save Albums in mysql database",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
