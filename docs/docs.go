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
        "/cages": {
            "get": {
                "description": "Retrieve a list of all cages.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of cages",
                "operationId": "get-cages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cage"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new cage in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new cage",
                "operationId": "create-cage",
                "parameters": [
                    {
                        "description": "Cage object",
                        "name": "cage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
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
        "/cages/{cage_id}/dinosaurs/{dinosaur_id}": {
            "post": {
                "description": "Add a dinosaur to a cage with specific checks.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a dinosaur to a cage.",
                "operationId": "add-dinosaur-to-cage",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cage ID",
                        "name": "cage_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Dinosaur ID",
                        "name": "dinosaur_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated cage",
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
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
        },
        "/cages/{id}": {
            "get": {
                "description": "Retrieve a specific cage based on its ID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a specific cage by ID",
                "operationId": "get-cage-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cage ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cage not found",
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
            },
            "put": {
                "description": "Update an existing cage in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an existing cage",
                "operationId": "update-cage",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cage ID to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Cage object",
                        "name": "cage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cage Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing cage by ID from the database.",
                "summary": "Delete an existing cage",
                "operationId": "delete-cage",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cage ID to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cage deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cage Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
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
        "models.Cage": {
            "type": "object",
            "properties": {
                "current_dinosaurs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Dinosaur"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "max_capacity": {
                    "type": "integer"
                },
                "power_status": {
                    "$ref": "#/definitions/models.PowerStatus"
                }
            }
        },
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
        },
        "models.PowerStatus": {
            "type": "string",
            "enum": [
                "DOWN",
                "ACTIVE",
                "DOWN"
            ],
            "x-enum-varnames": [
                "DefaultPowerStatus",
                "PowerStatusActive",
                "PowerStatusDown"
            ]
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
