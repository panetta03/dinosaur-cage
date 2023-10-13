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
        "/dinosaurs": {
            "post": {
                "description": "Create a new dinosaur",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a dinosaur",
                "operationId": "create-dinosaur",
                "parameters": [
                    {
                        "description": "Dinosaur object",
                        "name": "dinosaur",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Dinosaur"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Dinosaur"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dinosaurs/{id}": {
            "get": {
                "description": "Get a dinosaur's details by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a dinosaur by ID",
                "operationId": "get-dinosaur",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dinosaur ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dinosaur"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Diet": {
            "type": "string",
            "enum": [
                "carnivore",
                "herbivore"
            ],
            "x-enum-varnames": [
                "Carnivore",
                "Herbivore"
            ]
        },
        "models.Dinosaur": {
            "type": "object",
            "properties": {
                "cage_id": {
                    "type": "integer"
                },
                "diet": {
                    "$ref": "#/definitions/models.Diet"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "species": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
